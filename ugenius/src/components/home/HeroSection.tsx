import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { ArrowRight } from "lucide-react";

const FloatingParticle = ({ className }: { className?: string }) => (
  <div
    className={`absolute w-2 h-2 rounded-full bg-gold/40 animate-pulse-soft ${className}`}
  />
);

const FloatingShape = ({ className, children }: { className?: string; children?: React.ReactNode }) => (
  <div className={`absolute opacity-20 ${className}`}>
    {children}
  </div>
);

const HeroSection = () => {
  return (
    <section className="relative min-h-screen flex items-center justify-center overflow-hidden">
      {/* Background */}
      <div className="absolute inset-0 bg-gradient-hero">
        <div className="absolute inset-0 bg-gradient-to-b from-background/60 via-background/80 to-background" />
      </div>

      {/* Floating Elements */}
      <FloatingShape className="top-1/4 left-[10%] animate-float">
        <div className="w-20 h-20 border-2 border-gold/30 rounded-full" />
      </FloatingShape>
      <FloatingShape className="top-1/3 right-[15%] animate-float-delayed">
        <div className="w-16 h-16 border-2 border-coral/30 rotate-45" />
      </FloatingShape>
      <FloatingShape className="bottom-1/3 left-[20%] animate-float">
        <div className="w-12 h-12 bg-gold/10 rounded-lg rotate-12" />
      </FloatingShape>
      
      {/* Particles */}
      <FloatingParticle className="top-[20%] left-[30%]" />
      <FloatingParticle className="top-[40%] right-[25%]" />
      <FloatingParticle className="bottom-[30%] left-[15%]" />
      <FloatingParticle className="bottom-[40%] right-[20%]" />
      <FloatingParticle className="top-[60%] left-[40%]" />

      {/* Content */}
      <div className="relative z-10 container mx-auto px-4 pt-24">
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
          {/* Left Column - Text */}
          <div className="text-center lg:text-left">
            

            {/* Heading */}
            <h1 className="font-display text-5xl md:text-6xl lg:text-7xl font-bold text-foreground mb-6 animate-fade-up leading-tight">
              Unlock Your{" "}
              <span className="text-gradient-gold">Academic</span>
              <br />
              Genius Within
            </h1>

            {/* Subheading */}
            <p className="text-lg md:text-xl text-muted-foreground max-w-2xl mx-auto lg:mx-0 mb-10 animate-fade-up-delayed leading-relaxed">
              Join U-Genius, the premier academic excellence club empowering students
              to achieve first-class grades through proven strategies, mentorship,
              and a supportive community.
            </p>

            {/* CTA Buttons */}
            <div className="flex flex-col sm:flex-row gap-4 justify-center lg:justify-start animate-fade-up-delayed">
              <Link to="/join">
                <Button variant="hero" size="xl" className="group">
                  Join U-Genius Today
                  <ArrowRight className="w-5 h-5 group-hover:translate-x-1 transition-transform" />
                </Button>
              </Link>
              <Link to="/about">
                <Button variant="heroOutline" size="xl">
                  Learn More
                </Button>
              </Link>
            </div>

            {/* Stats */}
            <div className="mt-16 grid grid-cols-3 gap-8 max-w-xl mx-auto lg:mx-0 animate-fade-up-delayed">
              {[
                { number: "500+", label: "Students" },
                { number: "95%", label: "Success Rate" },
                { number: "50+", label: "Mentors" },
              ].map((stat) => (
                <div key={stat.label} className="text-center lg:text-left">
                  <div className="font-display text-3xl md:text-4xl font-bold text-gold">
                    {stat.number}
                  </div>
                  <div className="text-sm text-muted-foreground mt-1">{stat.label}</div>
                </div>
              ))}
            </div>
          </div>

          {/* Right Column - Hero Image */}
          <div className="hidden lg:block animate-fade-up-delayed">
            <div className="relative">
              {/* Main Image */}
              <div className="relative rounded-3xl overflow-hidden shadow-2xl">
                <img
                  src="https://images.unsplash.com/photo-1541339907198-e08756dedf3f?w=800&h=600&fit=crop&q=80"
                  alt="Graduates celebrating at convocation ceremony"
                  className="w-full h-[500px] object-cover"
                />
                <div className="absolute inset-0 bg-gradient-to-t from-background/60 via-transparent to-transparent" />
              </div>

              {/* Floating Card 1 */}
              <div className="absolute -bottom-6 -left-6 bg-card p-4 rounded-2xl shadow-medium animate-float">
                <div className="flex items-center gap-3">
                  <div className="w-12 h-12 rounded-full bg-gold/20 flex items-center justify-center">
                    <span className="text-2xl">🎓</span>
                  </div>
                  <div>
                    <p className="font-semibold text-foreground">First Class</p>
                    <p className="text-sm text-muted-foreground">Achievement Unlocked</p>
                  </div>
                </div>
              </div>

              {/* Floating Card 2 */}
              <div className="absolute -top-4 -right-4 bg-card p-4 rounded-2xl shadow-medium animate-float-delayed">
                <div className="flex items-center gap-3">
                  <div className="w-12 h-12 rounded-full bg-coral/20 flex items-center justify-center">
                    <span className="text-2xl">📚</span>
                  </div>
                  <div>
                    <p className="font-semibold text-foreground">100+ Hours</p>
                    <p className="text-sm text-muted-foreground">Study Logged</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* Scroll Indicator */}
      <div className="absolute bottom-8 left-1/2 -translate-x-1/2 animate-bounce">
        <div className="w-6 h-10 rounded-full border-2 border-muted-foreground/30 flex items-start justify-center p-1">
          <div className="w-1.5 h-3 rounded-full bg-gold animate-pulse" />
        </div>
      </div>
    </section>
  );
};

export default HeroSection;
