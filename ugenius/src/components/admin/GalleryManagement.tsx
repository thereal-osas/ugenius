import { useState, useEffect } from 'react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { Upload, X, Image as ImageIcon, Plus } from 'lucide-react';
import { toast } from '@/hooks/use-toast';

interface GalleryItem {
  id: number;
  title: string;
  image_url: string;
  caption: string;
  created_at: string;
}

interface GalleryManagementProps {
  onClose?: () => void;
}

export default function GalleryManagement({ onClose }: GalleryManagementProps) {
  const [galleryItems, setGalleryItems] = useState<GalleryItem[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [uploading, setUploading] = useState(false);
  const [showUploadForm, setShowUploadForm] = useState(false);
  const [uploadData, setUploadData] = useState({
    title: '',
    caption: '',
    imageFile: null as File | null,
    imagePreview: null as string | null
  });

  useEffect(() => {
    fetchGallery();
  }, []);

  const fetchGallery = async () => {
    try {
      const token = localStorage.getItem("access_token");
      const response = await fetch(`${import.meta.env.VITE_API_URL}/gallery`, {
        headers: {
          "Authorization": `Bearer ${token}`,
        },
      });
      
      if (response.ok) {
        const data = await response.json();
        console.log("Gallery data fetched:", data); // Debug log
        setGalleryItems(data.data || []);
      }
    } catch (error) {
      console.error("Failed to fetch gallery:", error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleImageSelect = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      // Create preview
      const reader = new FileReader();
      reader.onload = (e) => {
        setUploadData(prev => ({ 
          ...prev, 
          imageFile: file,
          imagePreview: e.target?.result as string
        }));
      };
      reader.readAsDataURL(file);
    }
  };

  const handleUpload = async () => {
    if (!uploadData.imageFile || !uploadData.title) {
      toast({
        title: "Error",
        description: "Please provide both an image and title",
        variant: "destructive",
      });
      return;
    }

    setUploading(true);
    
    try {
      const token = localStorage.getItem("access_token");
      
      // First upload the image
      const formData = new FormData();
      formData.append("image", uploadData.imageFile);

      const uploadResponse = await fetch(`${import.meta.env.VITE_API_URL}/admin/gallery/upload`, {
        method: "POST",
        headers: {
          "Authorization": `Bearer ${token}`,
        },
        body: formData,
      });

      if (!uploadResponse.ok) {
        const errorText = await uploadResponse.text();
        console.error("Upload failed:", errorText);
        throw new Error("Failed to upload image");
      }

      const uploadResult = await uploadResponse.json();
      const imageUrl = uploadResult.data.image_url;
      console.log("Upload successful, image URL:", imageUrl);

      // Then create the gallery item
      const createResponse = await fetch(`${import.meta.env.VITE_API_URL}/admin/gallery`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify({
          title: uploadData.title,
          image_url: imageUrl,
          caption: uploadData.caption || "",
        }),
      });

      if (createResponse.ok) {
        toast({
          title: "Success!",
          description: "Image uploaded successfully",
        });
        
        // Reset form
        setUploadData({ title: '', caption: '', imageFile: null, imagePreview: null });
        setShowUploadForm(false);
        
        // Wait a moment for file to be fully saved, then refresh gallery
        setTimeout(() => {
          fetchGallery();
        }, 1000); // 1 second delay
      } else {
        throw new Error("Failed to create gallery item");
      }
    } catch (error) {
      console.error("Upload error:", error);
      toast({
        title: "Error",
        description: "Failed to upload image",
        variant: "destructive",
      });
    } finally {
      setUploading(false);
    }
  };

  const handleDelete = async (id: number) => {
    if (!confirm("Are you sure you want to delete this image?")) return;

    try {
      const token = localStorage.getItem("access_token");
      const response = await fetch(`${import.meta.env.VITE_API_URL}/admin/gallery/${id}`, {
        method: "DELETE",
        headers: {
          "Authorization": `Bearer ${token}`,
        },
      });

      if (response.ok) {
        toast({
          title: "Success!",
          description: "Image deleted successfully",
        });
        setGalleryItems(galleryItems.filter(item => item.id !== id));
      } else {
        throw new Error("Failed to delete image");
      }
    } catch (err) {
      toast({
        title: "Error",
        description: "Failed to delete image",
        variant: "destructive",
      });
    }
  };

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex justify-between items-center">
        <div>
          <h2 className="text-2xl font-bold text-gray-900">Gallery Management</h2>
          <p className="text-gray-600">Upload and manage gallery images</p>
        </div>
        {onClose && (
          <Button variant="outline" onClick={onClose}>
            Back to Dashboard
          </Button>
        )}
      </div>

      {/* Upload Section */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Plus className="h-5 w-5" />
            Add New Image
          </CardTitle>
        </CardHeader>
        <CardContent>
          {showUploadForm ? (
            <div className="space-y-4">
              {/* Image Preview */}
              {uploadData.imagePreview && (
                <div className="mb-4">
                  <label className="block text-sm font-medium text-gray-700 mb-2">
                    Image Preview
                  </label>
                  <div className="aspect-square w-32 rounded-lg overflow-hidden bg-gray-100 border">
                    <img
                      src={uploadData.imagePreview}
                      alt="Preview"
                      className="w-full h-full object-cover"
                    />
                  </div>
                </div>
              )}
              
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  Image File
                </label>
                <input
                  type="file"
                  accept="image/*"
                  onChange={handleImageSelect}
                  className="block w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold file:bg-blue-50 file:text-blue-700 hover:file:bg-blue-100"
                />
              </div>
              
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  Title
                </label>
                <Input
                  value={uploadData.title}
                  onChange={(e) => setUploadData(prev => ({ ...prev, title: e.target.value }))}
                  placeholder="Enter image title"
                />
              </div>
              
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  Caption
                </label>
                <Textarea
                  value={uploadData.caption}
                  onChange={(e) => setUploadData(prev => ({ ...prev, caption: e.target.value }))}
                  placeholder="Enter image caption (optional)"
                  rows={3}
                />
              </div>
              
              <div className="flex gap-2">
                <Button 
                  onClick={handleUpload} 
                  disabled={uploading || !uploadData.imageFile || !uploadData.title}
                >
                  {uploading ? 'Uploading...' : 'Upload Image'}
                </Button>
                <Button 
                  variant="outline" 
                  onClick={() => {
                    setShowUploadForm(false);
                    setUploadData({ title: '', caption: '', imageFile: null, imagePreview: null });
                  }}
                >
                  Cancel
                </Button>
              </div>
            </div>
          ) : (
            <Button onClick={() => setShowUploadForm(true)} className="w-full">
              <Upload className="h-4 w-4 mr-2" />
              Upload New Image
            </Button>
          )}
        </CardContent>
      </Card>

      {/* Gallery Grid */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <ImageIcon className="h-5 w-5" />
            Uploaded Images ({galleryItems.length})
          </CardTitle>
        </CardHeader>
        <CardContent>
          {isLoading ? (
            <p className="text-center py-8">Loading gallery...</p>
          ) : galleryItems.length === 0 ? (
            <p className="text-center py-8 text-gray-500">
              No images uploaded yet. Upload your first image above.
            </p>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
              {galleryItems.map((item) => (
                <div key={item.id} className="relative group">
                  <div className="aspect-square rounded-lg overflow-hidden bg-gray-100">
                    <img
                      src={`${import.meta.env.VITE_API_URL.replace('/api/v1', '')}${item.image_url}`}
                      alt={item.title}
                      className="w-full h-full object-cover"
                    />
                  </div>
                  
                  {/* Delete Button */}
                  <button
                    onClick={() => handleDelete(item.id)}
                    className="absolute top-2 right-2 bg-red-500 text-white p-2 rounded-full opacity-0 group-hover:opacity-100 transition-opacity"
                  >
                    <X className="h-4 w-4" />
                  </button>
                  
                  {/* Image Info */}
                  <div className="mt-2">
                    <h3 className="font-medium text-gray-900 truncate">{item.title}</h3>
                    {item.caption && (
                      <p className="text-sm text-gray-600 line-clamp-2">{item.caption}</p>
                    )}
                  </div>
                </div>
              ))}
            </div>
          )}
        </CardContent>
      </Card>
    </div>
  );
}
