import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { GraduationCap, BookMarked, Users2, Lightbulb, ArrowRight } from "lucide-react";

const benefits = [
  {
    icon: GraduationCap,
    title: "Scholarships",
    description: "Access exclusive scholarship opportunities and financial support for exceptional performers.",
    color: "gold",
  },
  {
    icon: BookMarked,
    title: "Study Resources",
    description: "Get premium access to tutorials, past questions, and comprehensive study materials.",
    color: "coral",
  },
  {
    icon: Users2,
    title: "Mentorship",
    description: "Connect with experienced mentors and high-flyers who've achieved excellence in their fields.",
    color: "navy",
  },
  {
    icon: Lightbulb,
    title: "Skill Acquisition",
    description: "Develop essential skills beyond academics including technical and digital skills.",
    color: "gold",
  },
];

const BenefitsPreview = () => {
  return (
    <section className="py-24 bg-background relative">
      <div className="container mx-auto px-4">
        {/* Section Header */}
        <div className="flex flex-col lg:flex-row justify-between items-start lg:items-end gap-6 mb-16">
          <div>
            <span className="text-gold font-medium text-sm uppercase tracking-wider">
              Member Benefits
            </span>
            <h2 className="font-display text-4xl md:text-5xl font-bold text-foreground mt-4">
              What You'll Gain
            </h2>
          </div>
          <Link to="/benefits">
            <Button variant="ghost" className="group text-gold hover:text-gold hover:bg-gold/10">
              View All Benefits
              <ArrowRight className="w-4 h-4 ml-2 group-hover:translate-x-1 transition-transform" />
            </Button>
          </Link>
        </div>

        {/* Benefits Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          {benefits.map((benefit) => (
            <div
              key={benefit.title}
              className="group relative bg-card rounded-2xl p-8 shadow-soft hover:shadow-medium transition-all duration-500 hover:-translate-y-1 overflow-hidden"
            >
              {/* Background Gradient */}
              <div
                className={`absolute top-0 right-0 w-32 h-32 rounded-full blur-3xl opacity-20 transition-opacity duration-500 group-hover:opacity-40 ${
                  benefit.color === "gold"
                    ? "bg-gold"
                    : benefit.color === "coral"
                    ? "bg-coral"
                    : "bg-navy-light"
                }`}
              />

              <div className="relative z-10 flex gap-6">
                <div
                  className={`shrink-0 w-16 h-16 rounded-2xl flex items-center justify-center transition-all duration-300 group-hover:scale-110 ${
                    benefit.color === "gold"
                      ? "bg-gold/10 text-gold"
                      : benefit.color === "coral"
                      ? "bg-coral/20 text-coral-dark"
                      : "bg-navy/10 text-navy"
                  }`}
                >
                  <benefit.icon className="w-8 h-8" />
                </div>
                <div>
                  <h3 className="font-display text-xl font-semibold text-foreground mb-2">
                    {benefit.title}
                  </h3>
                  <p className="text-muted-foreground leading-relaxed">
                    {benefit.description}
                  </p>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
};

export default BenefitsPreview;
