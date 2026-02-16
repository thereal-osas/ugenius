import { Link } from "react-router-dom";
import Navbar from "@/components/layout/Navbar";
import Footer from "@/components/layout/Footer";
import { Button } from "@/components/ui/button";
import { CheckCircle, Home } from "lucide-react";

const ThankYou = () => {
  return (
    <div className="min-h-screen">
      <Navbar />
      <main>
        <section className="pt-32 pb-24 bg-gradient-hero relative overflow-hidden min-h-[80vh] flex items-center">
          <div className="absolute inset-0 opacity-30">
            <div className="absolute top-1/4 left-1/4 w-96 h-96 bg-gold/20 rounded-full blur-3xl" />
            <div className="absolute bottom-1/4 right-1/4 w-96 h-96 bg-coral/20 rounded-full blur-3xl" />
          </div>

          <div className="container mx-auto px-4 relative z-10">
            <div className="max-w-2xl mx-auto text-center">
              {/* Success Icon */}
              <div className="mb-8 inline-flex items-center justify-center w-24 h-24 rounded-full bg-gradient-gold shadow-gold">
                <CheckCircle className="w-12 h-12 text-primary" />
              </div>

              {/* Heading */}
              <h1 className="font-display text-4xl md:text-5xl font-bold text-foreground mb-6">
                Thank You for Registering!
              </h1>

              {/* Description */}
              <p className="text-lg text-muted-foreground leading-relaxed mb-8">
                You've successfully joined the U-Genius club. We're thrilled to have you 
                on board as we prepare to launch something truly amazing for students like you.
              </p>

              {/* What's Next Section */}
            
                

              {/* CTA */}
              <Link to="/">
                <Button variant="gold" size="lg">
                  <Home className="w-5 h-5 mr-2" />
                  Return to Home
                </Button>
              </Link>
            </div>
          </div>
        </section>
      </main>
      <Footer />
    </div>
  );
};

export default ThankYou;

