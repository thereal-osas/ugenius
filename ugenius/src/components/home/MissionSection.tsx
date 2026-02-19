import { Target, Users, BookOpen, Award } from "lucide-react";
import ScrollReveal from "@/components/ScrollReveal";

const features = [
  {
    icon: Target,
    title: "Cultivate Academic Excellence",
    description: "We equip students with the knowledge, discipline, and support needed to develop strong learning habits.",
  },
  {
    icon: Users,
    title: "Optimize Leadership Capacity",
    description: "We nurture confident, visionary students who are prepared to lead themselves and others with purpose and integrity.",
  },
  {
    icon: BookOpen,
    title: "Refine Character & Personal Growth",
    description: "We develop resilient, value-driven individuals who grow in confidence, mindset, and emotional intelligence.",
  },
  {
    icon: Award,
    title: "Empower Community Impact",
    description: "We inspire students to collaborate, serve, and create meaningful change within their campuses and beyond.",
  },
];

const MissionSection = () => {
  return (
    <section className="py-24 bg-cream-dark relative overflow-hidden">
      {/* Background Pattern */}
      <div className="absolute inset-0 opacity-30">
        <div className="absolute top-0 left-0 w-full h-full bg-gradient-to-br from-gold-100/20 via-transparent to-purple-100/20" />
      </div>

      <div className="container mx-auto px-4 relative z-10">
        {/* Section Header */}
        <ScrollReveal direction="up" delay={0.2}>
          <div className="text-center mb-16">
            <h2 className="font-display text-4xl md:text-5xl font-bold text-foreground mb-6">
              Our <span className="text-gradient-gold">Mission</span>
            </h2>
            <p className="text-lg md:text-xl text-muted-foreground max-w-3xl mx-auto leading-relaxed">
              Maintaining Academic Excellence, Inspiring Leadership
            </p>
          </div>
        </ScrollReveal>

        {/* Image + Text Section */}
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-16 items-center mb-20">
          {/* Image */}
          <div className="relative">
            <div className="rounded-3xl overflow-hidden shadow-xl">
              <img
                src="https://images.unsplash.com/photo-1456513080510-7bf3a84b82f8?w=800&h=500&fit=crop&q=80"
                alt="Student studying with books and notes"
                className="w-full h-[400px] object-cover"
              />
            </div>
            {/* Decorative element */}
            <div className="absolute -bottom-4 -right-4 w-24 h-24 bg-gold/20 rounded-2xl -z-10" />
            <div className="absolute -top-4 -left-4 w-16 h-16 bg-coral/20 rounded-xl -z-10" />
          </div>

          {/* Text */}
          <div className="lg:pl-16">
      
            <p className="text-lg md:text-xl text-muted-foreground max-w-3xl mx-auto leading-relaxed">
              We exist to raise a generation of scholars who not only succeed in the classroom, but also lead with character, vision, and impact. Our goal is to shape well-rounded individuals who are academically strong, emotionally grounded, and purpose-driven, prepared to stand out in any environment they find themselves in. <br/><br/> Through tutorials, mentorship, community, and skill-building initiatives, we equip students to think critically, grow confidently, and influence their world positively.
            </p>
          </div>
        </div>

<h3 className="font-semibold text-2xl text-center pb-8 uppercase tracking-wider">U-GENIUS <span className="text-gold font-bold">C.O.R.E</span> </h3>
        {/* Features Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
          {features.map((feature, index) => (
            <ScrollReveal 
              key={feature.title} 
              direction="up" 
              delay={0.3 + index * 0.1}
              className="group"
            >
              <div className="bg-white rounded-2xl p-8 shadow-soft hover:shadow-medium transition-all duration-300 group-hover:-translate-y-2">
                <div className="w-16 h-16 rounded-full bg-gradient-gold flex items-center justify-center mb-6 group-hover:scale-110 transition-transform">
                  <feature.icon className="w-8 h-8 text-primary" />
                </div>
                <h3 className="text-xl font-semibold text-foreground mb-4">
                  {feature.title}
                </h3>
                <p className="text-muted-foreground leading-relaxed">
                  {feature.description}
                </p>
              </div>
            </ScrollReveal>
          ))}
        </div>

        {/* Quote */}
        <ScrollReveal direction="up" delay={0.8}>
          <div className="mt-20 max-w-3xl mx-auto text-center">
            <blockquote className="relative">
             
              <p className="font-display text-2xl md:text-3xl text-foreground italic leading-relaxed relative z-10">
                "We are what we repeatedly do. Excellence, then, is not an act, but a habit."
              </p>
             
            </blockquote>
          </div>
        </ScrollReveal>
      </div>
    </section>
  );
};

export default MissionSection;
