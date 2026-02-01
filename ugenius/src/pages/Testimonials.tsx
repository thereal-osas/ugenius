import Navbar from "@/components/layout/Navbar";
import Footer from "@/components/layout/Footer";
import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Quote, ArrowRight } from "lucide-react";

const testimonials = [
  {
    id: 1,
    name: "Adaeze Okonkwo",
    role: "Computer Science Graduate",
    cgpa: "4.78",
    quote: "U-Genius transformed my approach to studying. From struggling to maintain a 3.0 GPA, I graduated with a First Class. The mentorship program was life-changing. My mentor helped me develop study strategies that actually worked for my learning style.",
    avatar: "AO",
  },
  {
    id: 2,
    name: "Emmanuel Nwachukwu",
    role: "Engineering, 400 Level",
    cgpa: "4.65",
    quote: "The Auto-Bio Jacking methodology taught me how to learn effectively. My grades improved dramatically from a 2.8 GPA to consistently above 4.5. I now mentor other students in the club and watch them achieve similar transformations.",
    avatar: "EN",
  },
  {
    id: 3,
    name: "Fatima Ibrahim",
    role: "Medicine, 500 Level",
    cgpa: "4.82",
    quote: "The study materials and support system at U-Genius are unmatched. I secured a scholarship through the connections I made here. The video tutorials for complex medical subjects saved me countless hours of confusion.",
    avatar: "FI",
  },
  {
    id: 4,
    name: "Chukwuemeka Obi",
    role: "Law Graduate",
    cgpa: "4.71",
    quote: "Joining U-Genius was the best decision of my university career. The time management techniques I learned helped me balance academics, moot court, and extra-curricular activities while still graduating with First Class honors.",
    avatar: "CO",
  },
  {
    id: 5,
    name: "Aisha Mohammed",
    role: "Biochemistry, 300 Level",
    cgpa: "4.55",
    quote: "As someone who struggled with science courses, U-Genius gave me the tools and confidence to excel. The peer study groups and past questions bank were invaluable. I went from barely passing to topping my class.",
    avatar: "AM",
  },
  {
    id: 6,
    name: "David Adeleke",
    role: "Economics Graduate",
    cgpa: "4.69",
    quote: "The mentorship I received at U-Genius extended beyond academics. My mentor helped me develop professional skills that landed me a job before graduation. The holistic approach to student development is what sets this club apart.",
    avatar: "DA",
  },
];

const Testimonials = () => {
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
                Success Stories
              </span>
              <h1 className="font-display text-5xl md:text-6xl font-bold text-foreground mt-4 mb-6">
                Voices of Excellence
              </h1>
              <p className="text-lg text-muted-foreground leading-relaxed">
                Hear from students who've transformed their academic journeys 
                through U-Genius and the Auto-Bio Jacking methodology.
              </p>
            </div>
          </div>
        </section>

        {/* Testimonials Grid */}
        <section className="py-24 bg-background">
          <div className="container mx-auto px-4">
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
              {testimonials.map((testimonial, index) => (
                <div
                  key={testimonial.id}
                  className={`group relative bg-card rounded-2xl p-8 shadow-soft hover:shadow-medium transition-all duration-500 hover:-translate-y-2 ${
                    index === 0 || index === 5 ? "md:col-span-2 lg:col-span-1" : ""
                  }`}
                >
                  {/* Quote Icon */}
                  <Quote className="absolute top-6 right-6 w-10 h-10 text-gold/20 group-hover:text-gold/30 transition-colors" />

                  {/* Content */}
                  <div className="relative z-10">
                    <p className="text-foreground leading-relaxed mb-6 italic">
                      "{testimonial.quote}"
                    </p>

                    <div className="flex items-center justify-between">
                      <div className="flex items-center gap-4">
                        <div className="w-14 h-14 rounded-full bg-gradient-gold flex items-center justify-center text-primary font-display font-bold text-lg">
                          {testimonial.avatar}
                        </div>
                        <div>
                          <div className="font-semibold text-foreground">
                            {testimonial.name}
                          </div>
                          <div className="text-muted-foreground text-sm">
                            {testimonial.role}
                          </div>
                        </div>
                      </div>
                      <div className="text-right">
                        <div className="text-gold font-display font-bold text-xl">
                          {testimonial.cgpa}
                        </div>
                        <div className="text-muted-foreground text-xs">CGPA</div>
                      </div>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </section>

        {/* Stats Section */}
        <section className="py-20 bg-gradient-navy text-primary-foreground">
          <div className="container mx-auto px-4">
            <div className="grid grid-cols-2 md:grid-cols-4 gap-8">
              {[
                { number: "95%", label: "Improvement Rate" },
                { number: "500+", label: "Active Members" },
                { number: "4.5+", label: "Average CGPA" },
                { number: "150+", label: "First Class Graduates" },
              ].map((stat) => (
                <div key={stat.label} className="text-center">
                  <div className="font-display text-4xl md:text-5xl font-bold text-gold">
                    {stat.number}
                  </div>
                  <div className="text-primary-foreground/70 mt-2">{stat.label}</div>
                </div>
              ))}
            </div>
          </div>
        </section>

        {/* CTA */}
        <section className="py-24 bg-background">
          <div className="container mx-auto px-4">
            <div className="max-w-3xl mx-auto text-center">
              <h2 className="font-display text-4xl md:text-5xl font-bold text-foreground mb-6">
                Write Your Success Story
              </h2>
              <p className="text-lg text-muted-foreground mb-8">
                Join U-Genius today and become the next success story. 
                Your journey to academic excellence starts here.
              </p>
              <Link to="/join">
                <Button variant="hero" size="xl" className="group">
                  Start Your Journey
                  <ArrowRight className="w-5 h-5 group-hover:translate-x-1 transition-transform" />
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

export default Testimonials;
