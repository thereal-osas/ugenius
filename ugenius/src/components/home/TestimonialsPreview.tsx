import { useState } from "react";
import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { ChevronLeft, ChevronRight, Quote, ArrowRight } from "lucide-react";

const testimonials = [
  {
    id: 1,
    name: "Adaeze Okonkwo",
    role: "Computer Science, Final Year",
    quote: "U-Genius transformed my approach to studying. From struggling to maintain a 3.0 GPA, I graduated with a First Class. The mentorship program was life-changing.",
    avatar: "AO",
  },
  {
    id: 2,
    name: "Emmanuel Nwachukwu",
    role: "Engineering, 400 Level",
    quote: "The Auto-Bio Jacking methodology taught me how to learn effectively. My grades improved dramatically, and I now mentor other students in the club.",
    avatar: "EN",
  },
  {
    id: 3,
    name: "Fatima Ibrahim",
    role: "Medicine, 500 Level",
    quote: "The study materials and support system at U-Genius are unmatched. I secured a scholarship through the connections I made here.",
    avatar: "FI",
  },
];

const TestimonialsPreview = () => {
  const [activeIndex, setActiveIndex] = useState(0);

  const nextTestimonial = () => {
    setActiveIndex((prev) => (prev + 1) % testimonials.length);
  };

  const prevTestimonial = () => {
    setActiveIndex((prev) => (prev - 1 + testimonials.length) % testimonials.length);
  };

  return (
    <section className="py-24 bg-gradient-navy text-primary-foreground relative overflow-hidden">
      {/* Background Elements */}
      <div className="absolute inset-0">
        <div className="absolute top-1/4 left-1/4 w-64 h-64 bg-gold/5 rounded-full blur-3xl" />
        <div className="absolute bottom-1/4 right-1/4 w-64 h-64 bg-coral/5 rounded-full blur-3xl" />
      </div>

      <div className="container mx-auto px-4 relative z-10">
        {/* Section Header */}
        <div className="flex flex-col lg:flex-row justify-between items-start lg:items-end gap-6 mb-16">
          <div>
            <span className="text-gold font-medium text-sm uppercase tracking-wider">
              Success Stories
            </span>
            <h2 className="font-display text-4xl md:text-5xl font-bold mt-4">
              What Our Members Say
            </h2>
          </div>
          <Link to="/testimonials">
            <Button variant="ghost" className="text-gold hover:text-gold hover:bg-gold/10 group">
              View All Stories
              <ArrowRight className="w-4 h-4 ml-2 group-hover:translate-x-1 transition-transform" />
            </Button>
          </Link>
        </div>

        {/* Testimonial Slider */}
        <div className="max-w-4xl mx-auto">
          <div className="relative">
            {/* Quote Icon */}
            <Quote className="absolute -top-4 -left-4 w-16 h-16 text-gold/20" />

            {/* Testimonial Content */}
            <div className="bg-primary-foreground/5 backdrop-blur-sm rounded-3xl p-8 md:p-12 border border-primary-foreground/10">
              <p className="text-xl md:text-2xl leading-relaxed mb-8 font-light">
                "{testimonials[activeIndex].quote}"
              </p>

              <div className="flex items-center gap-4">
                <div className="w-14 h-14 rounded-full bg-gold flex items-center justify-center text-primary font-display font-bold text-lg">
                  {testimonials[activeIndex].avatar}
                </div>
                <div>
                  <div className="font-semibold text-lg">
                    {testimonials[activeIndex].name}
                  </div>
                  <div className="text-primary-foreground/70 text-sm">
                    {testimonials[activeIndex].role}
                  </div>
                </div>
              </div>
            </div>

            {/* Navigation */}
            <div className="flex items-center justify-center gap-4 mt-8">
              <button
                onClick={prevTestimonial}
                className="w-12 h-12 rounded-full border border-primary-foreground/20 flex items-center justify-center hover:bg-primary-foreground/10 transition-colors"
              >
                <ChevronLeft className="w-5 h-5" />
              </button>

              <div className="flex gap-2">
                {testimonials.map((_, index) => (
                  <button
                    key={index}
                    onClick={() => setActiveIndex(index)}
                    className={`w-2 h-2 rounded-full transition-all duration-300 ${
                      index === activeIndex
                        ? "bg-gold w-8"
                        : "bg-primary-foreground/30 hover:bg-primary-foreground/50"
                    }`}
                  />
                ))}
              </div>

              <button
                onClick={nextTestimonial}
                className="w-12 h-12 rounded-full border border-primary-foreground/20 flex items-center justify-center hover:bg-primary-foreground/10 transition-colors"
              >
                <ChevronRight className="w-5 h-5" />
              </button>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default TestimonialsPreview;
