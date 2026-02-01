import Navbar from "@/components/layout/Navbar";
import Footer from "@/components/layout/Footer";
import { Target, Eye, Heart, Zap } from "lucide-react";

const values = [
  {
    icon: Target,
    title: "Excellence",
    description: "We pursue the highest standards in everything we do, inspiring students to aim for first-class achievement.",
  },
  {
    icon: Eye,
    title: "Vision",
    description: "We see the potential in every student and work to help them realize their academic dreams.",
  },
  {
    icon: Heart,
    title: "Community",
    description: "We believe in the power of togetherness, supporting each other through the academic journey.",
  },
  {
    icon: Zap,
    title: "Transformation",
    description: "We are committed to life-changing outcomes through proven methodologies and dedication.",
  },
];

const About = () => {
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
            <div className="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
              <div className="text-center lg:text-left">
                <span className="text-gold font-medium text-sm uppercase tracking-wider">
                  About U-Genius
                </span>
                <h1 className="font-display text-5xl md:text-6xl font-bold text-foreground mt-4 mb-6">
                  Our Story & Vision
                </h1>
                <p className="text-lg text-muted-foreground leading-relaxed">
                  U-Genius was born from a simple belief: every student has the potential
                  for academic excellence when given the right tools, guidance, and community.
                </p>
              </div>

              {/* Hero Image */}
              <div className="hidden lg:block">
                <div className="relative">
                  <div className="rounded-3xl overflow-hidden shadow-2xl">
                    <img
                      src="https://images.unsplash.com/photo-1541339907198-e08756dedf3f?w=800&h=450&fit=crop&q=80"
                      alt="Graduate in cap and gown at graduation ceremony"
                      className="w-full h-[350px] object-cover"
                    />
                  </div>
                  <div className="absolute -bottom-4 -left-4 w-20 h-20 bg-gold/30 rounded-2xl -z-10" />
                </div>
              </div>
            </div>
          </div>
        </section>

        {/* Mission & Vision */}
        <section className="py-24 bg-background">
          <div className="container mx-auto px-4">
            <div className="grid grid-cols-1 lg:grid-cols-2 gap-16">
              {/* Our Story */}
              <div>
                <span className="text-gold font-medium text-sm uppercase tracking-wider">
                  Our Story
                </span>
                <h2 className="font-display text-3xl md:text-4xl font-bold text-foreground mt-4 mb-6">
                  The Birth of U-Genius
                </h2>
                <div className="space-y-4 text-muted-foreground leading-relaxed">
                  <p>
                    U-Genius emerged from the transformative teachings of Dr. Isaiah Macwealth 
                    and his revolutionary Auto-Bio Jacking system. This methodology has helped 
                    countless students break through academic barriers and achieve excellence.
                  </p>
                  <p>
                    Founded by a group of passionate students who experienced firsthand the 
                    power of these principles, U-Genius became a beacon for those aspiring 
                    to become first-class graduates.
                  </p>
                  <p>
                    Today, we are a thriving community of learners, mentors, and achievers, 
                    all united by the common goal of academic excellence and personal growth.
                  </p>
                </div>
              </div>

              {/* Auto-Bio Jacking */}
              <div className="bg-card rounded-3xl p-8 shadow-soft">
                <span className="text-coral-dark font-medium text-sm uppercase tracking-wider">
                  The Methodology
                </span>
                <h2 className="font-display text-3xl md:text-4xl font-bold text-foreground mt-4 mb-6">
                  Auto-Bio Jacking
                </h2>
                <div className="space-y-4 text-muted-foreground leading-relaxed">
                  <p>
                    Auto-Bio Jacking is a unique approach to learning and personal development 
                    that combines self-discovery, strategic thinking, and proven study techniques.
                  </p>
                  <p>
                    At its core, the system teaches students to:
                  </p>
                  <ul className="space-y-3">
                    {[
                      "Understand their unique learning patterns",
                      "Develop effective study strategies",
                      "Build mental resilience and focus",
                      "Create sustainable habits for success",
                      "Connect with mentors and peers for growth",
                    ].map((item) => (
                      <li key={item} className="flex items-start gap-3">
                        <div className="w-1.5 h-1.5 rounded-full bg-gold mt-2 shrink-0" />
                        <span>{item}</span>
                      </li>
                    ))}
                  </ul>
                </div>
              </div>
            </div>
          </div>
        </section>

        {/* Core Values */}
        <section className="py-24 bg-cream-dark">
          <div className="container mx-auto px-4">
            <div className="text-center mb-16">
              <span className="text-gold font-medium text-sm uppercase tracking-wider">
                What We Stand For
              </span>
              <h2 className="font-display text-4xl md:text-5xl font-bold text-foreground mt-4">
                Our Core Values
              </h2>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
              {values.map((value) => (
                <div
                  key={value.title}
                  className="bg-card rounded-2xl p-8 text-center shadow-soft hover:shadow-medium transition-all duration-500 hover:-translate-y-2"
                >
                  <div className="w-16 h-16 rounded-2xl bg-gold/10 flex items-center justify-center mx-auto mb-6">
                    <value.icon className="w-8 h-8 text-gold" />
                  </div>
                  <h3 className="font-display text-xl font-semibold text-foreground mb-3">
                    {value.title}
                  </h3>
                  <p className="text-muted-foreground text-sm leading-relaxed">
                    {value.description}
                  </p>
                </div>
              ))}
            </div>
          </div>
        </section>

        {/* Vision Statement */}
        <section className="py-24 bg-gradient-navy text-primary-foreground">
          <div className="container mx-auto px-4">
            <div className="max-w-4xl mx-auto text-center">
              <span className="text-gold font-medium text-sm uppercase tracking-wider">
                Our Vision
              </span>
              <h2 className="font-display text-4xl md:text-5xl font-bold mt-4 mb-8">
                Building Tomorrow's Leaders
              </h2>
              <p className="text-xl leading-relaxed text-primary-foreground/80">
                We envision a world where every student has access to the tools, mentorship, 
                and community they need to achieve academic excellence. Through U-Genius, 
                we're not just creating first-class graduates—we're nurturing future leaders, 
                innovators, and changemakers who will transform their communities and beyond.
              </p>
            </div>
          </div>
        </section>
      </main>
      <Footer />
    </div>
  );
};

export default About;
