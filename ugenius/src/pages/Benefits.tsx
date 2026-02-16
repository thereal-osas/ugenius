import Navbar from "@/components/layout/Navbar";
import Footer from "@/components/layout/Footer";
import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import {
  GraduationCap,
  BookMarked,
  Users2,
  Lightbulb,
  Award,
  Calendar,
  ArrowRight,
} from "lucide-react";

const benefits = [
  {
    icon: GraduationCap,
    title: "Scholarship Opportunities",
    description: "Access funding support, academic grants, and exclusive opportunities that reward excellence and reduce financial barriers to educational advancement.",
    highlights: ["Merit-based awards", "Need-based assistance", "Research grants"],
  },
  {
    icon: BookMarked,
    title: "Tutorials & Premium Study Materials",
    description: "Get structured tutorials, curated resources, and premium study materials designed to simplify learning, improve understanding, and boost academic performance.",
    highlights: ["Past questions bank", "Solved examples", "Course summaries"],
  },
  {
    icon: Users2,
    title: "Mentorship Programs",
    description: "Connect with experienced mentors who provide guidance, accountability, and career insight to help you make informed academic and life decisions.",
    highlights: ["1-on-1 sessions", "Group mentoring", "Career guidance"],
  },
  {
    icon: Lightbulb,
    title: "Skill Development",
    description: "Build essential leadership, digital, and technical skills that prepare you for real-world challenges and future career success.",
    highlights: ["Leadership training", "Time management", "Communication skills"],
  },
  {
    icon: Calendar,
    title: "Events & Workshops",
    description: "Participate in impactful events and practical workshops that expand knowledge, encourage networking, and inspire personal and professional growth.",
    highlights: ["Weekly seminars", "Study groups", "Academic conferences"],
  },
  {
    icon: Award,
    title: "Recognition & Awards",
    description: "Earn recognition for excellence, leadership, and commitment through awards that motivate achievement and celebrate outstanding student contributions.",
    highlights: ["Member spotlight", "Certificates", "Annual awards"],
  },
];

const Benefits = () => {
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
                Member Benefits
              </span>
              <h1 className="font-display text-5xl md:text-6xl font-bold text-foreground mt-4 mb-6">
                Everything You Need to Excel
              </h1>
              <p className="text-lg text-muted-foreground leading-relaxed">
                As a U-Genius member, you gain access to a comprehensive suite of resources, 
                support systems, and opportunities designed to propel you towards academic excellence and leadership.
              </p>
            </div>
          </div>
        </section>

        {/* Benefits Grid */}
        <section className="py-24 bg-background">
          <div className="container mx-auto px-4">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
              {benefits.map((benefit) => (
                <div
                  key={benefit.title}
                  className="group bg-card rounded-2xl p-8 shadow-soft hover:shadow-medium transition-all duration-500 hover:-translate-y-1"
                >
                  <div className="flex gap-6">
                    <div className="shrink-0 w-16 h-16 rounded-2xl bg-gold/10 flex items-center justify-center group-hover:bg-gold/20 group-hover:scale-110 transition-all duration-300">
                      <benefit.icon className="w-8 h-8 text-gold" />
                    </div>
                    <div className="flex-1">
                      <h3 className="font-display text-xl font-semibold text-foreground mb-3">
                        {benefit.title}
                      </h3>
                      <p className="text-muted-foreground leading-relaxed mb-4">
                        {benefit.description}
                      </p>
                      <div className="flex flex-wrap gap-2">
                        {benefit.highlights.map((highlight) => (
                          <span
                            key={highlight}
                            className="px-3 py-1 text-xs font-medium rounded-full bg-gold/10 text-gold"
                          >
                            {highlight}
                          </span>
                        ))}
                      </div>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </section>

        {/* CTA */}
        <section className="py-24 bg-cream-dark">
          <div className="container mx-auto px-4">
            <div className="max-w-3xl mx-auto text-center">
              <h2 className="font-display text-4xl md:text-5xl font-bold text-foreground mb-6">
                Ready to Access These Benefits?
              </h2>
              <p className="text-lg text-muted-foreground mb-8">
                Join U-Genius today and start your journey towards academic excellence. 
                All benefits are available immediately upon membership.
              </p>
              <Link to="/join">
                <Button variant="hero" size="xl" className="group">
                  Join U-Genius Now
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

export default Benefits;
