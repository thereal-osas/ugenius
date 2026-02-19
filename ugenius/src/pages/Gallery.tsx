import { useState, useEffect, useCallback } from "react";
import Navbar from "@/components/layout/Navbar";
import Footer from "@/components/layout/Footer";
import { Card, CardContent } from "@/components/ui/card";
import { Image as ImageIcon, X, ChevronLeft, ChevronRight } from "lucide-react";
import { useAuth } from "@/contexts/AuthContext";

interface GalleryItem {
  id: number;
  title: string;
  image_url: string;
  caption: string;
  created_at: string;
}

const Gallery = () => {
  const { user, isAuthenticated } = useAuth();
  const [galleryItems, setGalleryItems] = useState<GalleryItem[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [selectedImage, setSelectedImage] = useState<GalleryItem | null>(null);
  const [currentIndex, setCurrentIndex] = useState(0);

  const isAdmin = isAuthenticated && user?.role === 'super_admin';

  useEffect(() => {
    fetchGallery();
  }, [user]);

  const fetchGallery = async () => {
    try {
      const response = await fetch(`${import.meta.env.VITE_API_URL}/gallery`);
      
      if (response.ok) {
        const data = await response.json();
        console.log("Gallery data:", data); // Debug log
        setGalleryItems(data.data || []);
      }
    } catch (error) {
      console.error("Failed to fetch gallery:", error);
    } finally {
      setIsLoading(false);
    }
  };

  const openModal = (item: GalleryItem, index: number) => {
    setSelectedImage(item);
    setCurrentIndex(index);
  };

  const closeModal = useCallback(() => {
    setSelectedImage(null);
  }, []);

  const navigatePrevious = useCallback(() => {
    const newIndex = currentIndex === 0 ? galleryItems.length - 1 : currentIndex - 1;
    setCurrentIndex(newIndex);
    setSelectedImage(galleryItems[newIndex]);
  }, [currentIndex, galleryItems]);

  const navigateNext = useCallback(() => {
    const newIndex = currentIndex === galleryItems.length - 1 ? 0 : currentIndex + 1;
    setCurrentIndex(newIndex);
    setSelectedImage(galleryItems[newIndex]);
  }, [currentIndex, galleryItems]);

  useEffect(() => {
    const handleGlobalKeyDown = (e: KeyboardEvent) => {
      if (!selectedImage) return;
      
      if (e.key === 'Escape') closeModal();
      if (e.key === 'ArrowLeft') navigatePrevious();
      if (e.key === 'ArrowRight') navigateNext();
    };
    
    if (selectedImage) {
      document.addEventListener('keydown', handleGlobalKeyDown);
      document.body.style.overflow = 'hidden';
    } else {
      document.removeEventListener('keydown', handleGlobalKeyDown);
      document.body.style.overflow = 'unset';
    }
    
    return () => {
      document.removeEventListener('keydown', handleGlobalKeyDown);
      document.body.style.overflow = 'unset';
    };
  }, [selectedImage, currentIndex, closeModal, navigatePrevious, navigateNext]);

  if (isLoading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-gold"></div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-background">
      <Navbar />
      
      <main className="container mx-auto px-4 py-8">
        <div className="mt-20 mb-8">
          <p className="text-lg text-muted-foreground mb-8">
            Explore moments from our academic journey and community events.
          </p>
        </div>

        {/* Gallery Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
          {galleryItems.map((item, index) => (
            <Card 
              key={item.id} 
              className="group overflow-hidden hover:shadow-lg transition-shadow duration-300 cursor-pointer"
              onClick={() => openModal(item, index)}
            >
              <CardContent className="p-0">
                <div className="relative">
                  <img
                    src={`${import.meta.env.VITE_API_URL.replace('/api/v1', '')}${item.image_url}`}
                    alt={item.title}
                    className="w-full h-64 object-cover group-hover:scale-105 transition-transform duration-300"
                    onError={() => {
                      console.error("Image failed to load:", {
                        fullUrl: `${import.meta.env.VITE_API_URL.replace('/api/v1', '')}${item.image_url}`,
                        itemUrl: item.image_url,
                        envUrl: import.meta.env.VITE_API_URL
                      });
                    }}
                    onLoad={() => {
                      console.log("Image loaded successfully:", `${import.meta.env.VITE_API_URL.replace('/api/v1', '')}${item.image_url}`);
                    }}
                  />
                  <div className="absolute inset-0 bg-black/0 group-hover:bg-black/20 transition-colors duration-300 flex items-center justify-center">
                    <div className="opacity-0 group-hover:opacity-100 transition-opacity duration-300">
                      <div className="bg-white/90 rounded-full p-2">
                        <ImageIcon className="w-6 h-6 text-gray-800" />
                      </div>
                    </div>
                  </div>
                </div>
                
                <div className="p-4">
                  {/* <h3 className="font-semibold text-lg mb-2">{item.title}</h3>
                  {item.caption && (
                    <p className="text-muted-foreground text-sm">{item.caption}</p>
                  )} */}
                  {/* <p className="text-xs text-muted-foreground mt-2">
                    {new Date(item.created_at).toLocaleDateString()}
                  </p> */}
                </div>
              </CardContent>
            </Card>
          ))}
        </div>

        {/* Empty State */}
        {galleryItems.length === 0 && (
          <div className="text-center py-12">
            <ImageIcon className="w-16 h-16 mx-auto text-muted-foreground/30 mb-4" />
            <h3 className="text-xl font-semibold text-muted-foreground mb-2">
              No images in gallery yet
            </h3>
            <p className="text-muted-foreground">
              {isAdmin 
                ? "Upload your first image to get started!"
                : "Check back soon for amazing moments from our community."
              }
            </p>
          </div>
        )}
      </main>

      <Footer />

      {/* Image Modal/Lightbox */}
      {selectedImage && (
        <div 
          className="fixed inset-0 z-50 flex items-center justify-center bg-black/90 p-4"
          onClick={closeModal}
        >
          <div 
            className="relative max-w-6xl max-h-full w-full h-full flex items-center justify-center"
            onClick={(e) => e.stopPropagation()}
          >
            {/* Close Button */}
            <button
              onClick={closeModal}
              className="absolute top-4 right-4 z-10 bg-white/10 hover:bg-white/20 text-white rounded-full p-2 transition-colors"
            >
              <X className="w-6 h-6" />
            </button>

            {/* Previous Button */}
            {galleryItems.length > 1 && (
              <button
                onClick={navigatePrevious}
                className="absolute left-4 z-10 bg-white/10 hover:bg-white/20 text-white rounded-full p-2 transition-colors"
              >
                <ChevronLeft className="w-6 h-6" />
              </button>
            )}

            {/* Next Button */}
            {galleryItems.length > 1 && (
              <button
                onClick={navigateNext}
                className="absolute right-4 z-10 bg-white/10 hover:bg-white/20 text-white rounded-full p-2 transition-colors"
              >
                <ChevronRight className="w-6 h-6" />
              </button>
            )}

            {/* Image */}
            <img
              src={`${import.meta.env.VITE_API_URL.replace('/api/v1', '')}${selectedImage.image_url}`}
              alt={selectedImage.title}
              className="max-w-full max-h-full object-contain"
            />

            {/* Image Info */}
            {selectedImage.title && (
              <div className="absolute bottom-4 left-4 right-4 text-center">
                <h3 className="text-white text-lg font-semibold">{selectedImage.title}</h3>
                {selectedImage.caption && (
                  <p className="text-white/80 text-sm mt-1">{selectedImage.caption}</p>
                )}
              </div>
            )}

            {/* Image Counter */}
            {galleryItems.length > 1 && (
              <div className="absolute top-4 left-4 text-white/80 text-sm">
                {currentIndex + 1} / {galleryItems.length}
              </div>
            )}
          </div>
        </div>
      )}
    </div>
  );
};

export default Gallery;
