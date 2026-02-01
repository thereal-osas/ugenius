import { Link } from "react-router-dom";
import Navbar from "@/components/layout/Navbar";
import Footer from "@/components/layout/Footer";
import { Button } from "@/components/ui/button";
import { CheckCircle, Mail, Bell, Home } from "lucide-react";

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
                You've successfully joined the U-Genius waitlist. We're thrilled to have you 
                on board as we prepare to launch something truly amazing for students like you.
              </p>

              {/* What's Next Section */}
              <div className="bg-card rounded-2xl p-8 mb-8 shadow-soft text-left">
                <h2 className="font-display text-xl font-semibold text-foreground mb-6 text-center">
                  What Happens Next?
                </h2>
                
                <div className="space-y-4">
                  <div className="flex items-start gap-4">
                    <div className="w-10 h-10 rounded-full bg-gold/10 flex items-center justify-center flex-shrink-0">
                      <Mail className="w-5 h-5 text-gold" />
                    </div>
                    <div>
                      <h3 className="font-medium text-foreground">Check Your Email</h3>
                      <p className="text-sm text-muted-foreground">
                        You'll receive a confirmation email shortly with more details about U-Genius.
                      </p>
                    </div>
                  </div>

                  <div className="flex items-start gap-4">
                    <div className="w-10 h-10 rounded-full bg-gold/10 flex items-center justify-center flex-shrink-0">
                      <Bell className="w-5 h-5 text-gold" />
                    </div>
                    <div>
                      <h3 className="font-medium text-foreground">Stay Tuned</h3>
                      <p className="text-sm text-muted-foreground">
                        We'll notify you when U-Genius launches at your campus with exclusive early access benefits.
                      </p>
                    </div>
                  </div>
                </div>
              </div>

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

