import { Target, Users, BookOpen, Award } from "lucide-react";

const features = [
  {
    icon: Target,
    title: "Focused Excellence",
    description: "We help students set clear academic goals and develop strategies to achieve first-class honors.",
  },
  {
    icon: Users,
    title: "Peer Community",
    description: "Connect with like-minded students who share your passion for academic achievement.",
  },
  {
    icon: BookOpen,
    title: "Quality Resources",
    description: "Access exclusive study materials, tutorials, and guides curated by top performers.",
  },
  {
    icon: Award,
    title: "Recognition",
    description: "Earn scholarships and recognition for your academic achievements and contributions.",
  },
];

const MissionSection = () => {
  return (
    <section className="py-24 bg-cream-dark relative overflow-hidden">
      {/* Background Pattern */}
      <div className="absolute inset-0 opacity-30">
        <div className="absolute top-0 left-0 w-96 h-96 bg-gold/10 rounded-full blur-3xl" />
        <div className="absolute bottom-0 right-0 w-96 h-96 bg-coral/10 rounded-full blur-3xl" />
      </div>

      <div className="container mx-auto px-4 relative z-10">
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

          {/* Text Content */}
          <div>
            <span className="text-gold font-medium text-sm uppercase tracking-wider">
              Our Mission
            </span>
            <h2 className="font-display text-4xl md:text-5xl font-bold text-foreground mt-4 mb-6">
              Empowering Academic Excellence
            </h2>
            <p className="text-muted-foreground text-lg leading-relaxed mb-6">
              U-Genius is more than a club—it's a transformative journey designed to unlock
              your full academic potential through the proven Auto-Bio Jacking methodology.
            </p>
            <p className="text-muted-foreground leading-relaxed">
              We believe every student has untapped genius waiting to be discovered. Through
              structured study habits, peer support, and expert guidance, we help you develop
              the mindset and skills needed to excel academically.
            </p>
          </div>
        </div>

        {/* Features Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
          {features.map((feature, index) => (
            <div
              key={feature.title}
              className="group bg-card rounded-2xl p-8 shadow-soft hover:shadow-medium transition-all duration-500 hover:-translate-y-2"
              style={{ animationDelay: `${index * 100}ms` }}
            >
              <div className="w-14 h-14 rounded-xl bg-gold/10 flex items-center justify-center mb-6 group-hover:bg-gold/20 group-hover:scale-110 transition-all duration-300">
                <feature.icon className="w-7 h-7 text-gold" />
              </div>
              <h3 className="font-display text-xl font-semibold text-foreground mb-3">
                {feature.title}
              </h3>
              <p className="text-muted-foreground leading-relaxed">
                {feature.description}
              </p>
            </div>
          ))}
        </div>

        {/* Quote */}
        <div className="mt-20 max-w-3xl mx-auto text-center">
          <blockquote className="relative">
            <div className="absolute -top-8 left-1/2 -translate-x-1/2 text-8xl text-gold/20 font-display">
              "
            </div>
            <p className="font-display text-2xl md:text-3xl text-foreground italic leading-relaxed relative z-10">
              Excellence is not a destination but a continuous journey of self-improvement 
              and dedication.
            </p>
            <footer className="mt-6">
              <span className="text-gold font-semibold">Dr. Isaiah Macwealth</span>
              <span className="text-muted-foreground"> — Creator of Auto-Bio Jacking</span>
            </footer>
          </blockquote>
        </div>
      </div>
    </section>
  );
};

export default MissionSection;
