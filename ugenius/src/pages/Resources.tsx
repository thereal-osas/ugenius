import Navbar from "@/components/layout/Navbar";
import Footer from "@/components/layout/Footer";
import { Button } from "@/components/ui/button";
import { FileText, Video, Download, BookOpen, ExternalLink } from "lucide-react";

const categories = [
  { id: "all", label: "All Resources" },
  { id: "pdf", label: "PDF Guides" },
  { id: "video", label: "Video Tutorials" },
  { id: "study", label: "Study Materials" },
];

const resources = [
  {
    id: 1,
    title: "Complete Guide to First Class Success",
    description: "A comprehensive PDF guide covering proven strategies for achieving first-class honors.",
    type: "pdf",
    icon: FileText,
    downloadable: true,
  },
  {
    id: 2,
    title: "Auto-Bio Jacking Fundamentals",
    description: "Video series introducing the core principles of Dr. Macwealth's methodology.",
    type: "video",
    icon: Video,
    downloadable: false,
  },
  {
    id: 3,
    title: "Time Management Mastery",
    description: "Learn how to effectively manage your study time for maximum productivity.",
    type: "pdf",
    icon: FileText,
    downloadable: true,
  },
  {
    id: 4,
    title: "Effective Note-Taking Techniques",
    description: "Video tutorial on advanced note-taking methods that improve retention.",
    type: "video",
    icon: Video,
    downloadable: false,
  },
  {
    id: 5,
    title: "Past Questions Compilation",
    description: "Extensive collection of past exam questions across various departments.",
    type: "study",
    icon: BookOpen,
    downloadable: true,
  },
  {
    id: 6,
    title: "Exam Preparation Blueprint",
    description: "Step-by-step guide to preparing for and excelling in examinations.",
    type: "pdf",
    icon: FileText,
    downloadable: true,
  },
  {
    id: 7,
    title: "Memory Enhancement Workshop",
    description: "Video workshop on techniques to improve memory and recall.",
    type: "video",
    icon: Video,
    downloadable: false,
  },
  {
    id: 8,
    title: "Course Registration Strategy Guide",
    description: "Tips and strategies for optimal course selection and registration.",
    type: "study",
    icon: BookOpen,
    downloadable: true,
  },
];

const Resources = () => {
  return (
    <div className="min-h-screen">
      <Navbar />
      <main>
        {/* Hero Section */}
        <section className="pt-32 pb-20 bg-gradient-hero relative overflow-hidden">
          <div className="absolute inset-0 opacity-30">
            <div className="absolute top-1/4 left-1/4 w-96 h-96 bg-gold/20 rounded-full blur-3xl" />
            <div className="absolute bottom-1/4 right-1/4 w-96 h-96 bg-coral/20 rounded-full blur-3xl" />
          </div>

          <div className="container mx-auto px-4 relative z-10">
            <div className="max-w-3xl mx-auto text-center">
              <span className="text-gold font-medium text-sm uppercase tracking-wider">
                Resource Center
              </span>
              <h1 className="font-display text-5xl md:text-6xl font-bold text-foreground mt-4 mb-6">
                Study Materials & Resources
              </h1>
              <p className="text-lg text-muted-foreground leading-relaxed">
                Access our comprehensive library of study materials, video tutorials, 
                and guides designed to accelerate your academic success.
              </p>
            </div>
          </div>
        </section>

        {/* Category Filter */}
        <section className="py-8 bg-background border-b border-border sticky top-16 z-30 backdrop-blur-md bg-background/95">
          <div className="container mx-auto px-4">
            <div className="flex flex-wrap gap-3 justify-center">
              {categories.map((category) => (
                <button
                  key={category.id}
                  className={`px-6 py-2 rounded-full text-sm font-medium transition-all duration-300 ${
                    category.id === "all"
                      ? "bg-gold text-primary shadow-gold"
                      : "bg-muted text-muted-foreground hover:bg-gold/10 hover:text-foreground"
                  }`}
                >
                  {category.label}
                </button>
              ))}
            </div>
          </div>
        </section>

        {/* Resources Grid */}
        <section className="py-16 bg-background">
          <div className="container mx-auto px-4">
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {resources.map((resource) => (
                <div
                  key={resource.id}
                  className="group bg-card rounded-2xl p-6 shadow-soft hover:shadow-medium transition-all duration-500 hover:-translate-y-1"
                >
                  <div className="flex items-start gap-4">
                    <div className={`shrink-0 w-12 h-12 rounded-xl flex items-center justify-center ${
                      resource.type === "pdf"
                        ? "bg-red-100 text-red-600"
                        : resource.type === "video"
                        ? "bg-blue-100 text-blue-600"
                        : "bg-gold/10 text-gold"
                    }`}>
                      <resource.icon className="w-6 h-6" />
                    </div>
                    <div className="flex-1">
                      <span className={`text-xs font-medium uppercase tracking-wider ${
                        resource.type === "pdf"
                          ? "text-red-600"
                          : resource.type === "video"
                          ? "text-blue-600"
                          : "text-gold"
                      }`}>
                        {resource.type}
                      </span>
                      <h3 className="font-display text-lg font-semibold text-foreground mt-1 mb-2 group-hover:text-gold transition-colors">
                        {resource.title}
                      </h3>
                      <p className="text-muted-foreground text-sm leading-relaxed mb-4">
                        {resource.description}
                      </p>
                      <Button
                        variant="ghost"
                        size="sm"
                        className="text-gold hover:text-gold hover:bg-gold/10"
                      >
                        {resource.downloadable ? (
                          <>
                            <Download className="w-4 h-4 mr-2" />
                            Download
                          </>
                        ) : (
                          <>
                            <ExternalLink className="w-4 h-4 mr-2" />
                            Watch
                          </>
                        )}
                      </Button>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </section>

        {/* Members Notice */}
        <section className="py-16 bg-cream-dark">
          <div className="container mx-auto px-4">
            <div className="max-w-2xl mx-auto text-center">
              <h2 className="font-display text-2xl font-semibold text-foreground mb-4">
                Access More Resources
              </h2>
              <p className="text-muted-foreground mb-6">
                Members get unlimited access to our complete resource library including 
                premium video courses and exclusive study materials.
              </p>
              <Button variant="gold" size="lg">
                Become a Member
              </Button>
            </div>
          </div>
        </section>
      </main>
      <Footer />
    </div>
  );
};

export default Resources;
