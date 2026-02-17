import { useState, useEffect } from "react";
import Navbar from "@/components/layout/Navbar";
import Footer from "@/components/layout/Footer";
import { Card, CardContent } from "@/components/ui/card";
import { Image as ImageIcon } from "lucide-react";
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
          {galleryItems.map((item) => (
            <Card key={item.id} className="group overflow-hidden hover:shadow-lg transition-shadow duration-300">
              <CardContent className="p-0">
                <div className="relative">
                  <img
                    src={`${import.meta.env.VITE_API_URL.replace('/api/v1', '')}${item.image_url}`}
                    alt={item.title}
                    className="w-full h-64 object-cover"
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
    </div>
  );
};

export default Gallery;
