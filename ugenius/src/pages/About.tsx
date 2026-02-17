import Navbar from "@/components/layout/Navbar";
import Footer from "@/components/layout/Footer";
import { Target, Eye, Heart, Zap } from "lucide-react";
import prophet from "@/assets/prophet.jpeg"

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
                 UGenius was created out of a deep desire to see students rise beyond limitations and unlock their full academic and leadership potential. What began as a simple initiative to support struggling students has grown into a movement committed to excellence, purpose, and impact.
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
                  UGenius was born on campus through firsthand encounters with students facing academic pressure, lack of guidance, and limited access to the right resources. Seeing the gap between potential and performance, UGenius emerged as a structured solution—combining academic support, mentorship, and leadership development into one empowering platform.
                  </p>
                  <p>
                    From these early experiences, we realized that success is not just about intelligence—it is about discipline, consistency, and the right systems. UGenius was created to provide students with the tools, structure, and encouragement they need to study daily, grow holistically, and achieve meaningful results.
                  </p>
                  <p>
                   Today, UGenius continues to expand, guided by the same vision: to raise a generation of scholars and leaders who are equipped to excel academically and make lasting impact in their communities.
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
                    Autobiojacking, also known as ABJ, is the heart of UGenius’ approach to academic success. It is a structured system of discipline designed to ensure that students study consistently every day, turning learning into a daily habit rather than a last-minute effort.</p>

<p>ABJ is more than just a schedule—it is a system of operation that aligns mindset, routine, and accountability. By embedding consistent study habits into students’ daily lives, ABJ helps them retain knowledge better, improve performance, and build the discipline essential for long-term academic achievement.</p>

<p>At UGenius, we believe that consistency is the key to success, and Autobiogaking provides the framework that turns potential into results. Through ABJ, students are empowered to take control of their learning, track progress, and steadily move toward academic excellence</p>

                  
                 
                  
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

        {/* Founder Section */}
        <section className="py-20 bg-background">
          <div className="container mx-auto px-4">
            <div className="max-w-6xl mx-auto">
              <div className="text-center mb-12">
                <span className="text-gold font-medium text-sm uppercase tracking-wider">
                  Meet Our Founder
                </span>
                <h2 className="font-display text-4xl md:text-5xl font-bold text-foreground mt-4 mb-6">
                  The Visionary Behind U-Genius
                </h2>
                <p className="text-lg text-muted-foreground max-w-3xl mx-auto">
                  A visionary Christian leader, mentor, author, and transformational education advocate.
                </p>
              </div>
              
              <div className="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center">
                <div className="order-2 lg:order-1">
                  <div className="space-y-6">
                    <div>
                      <h3 className="font-display text-2xl font-bold text-foreground mb-4">
                        Dr. Isaiah Macwealth
                      </h3>
                      <p className="text-gold font-medium mb-4">
                        Founder
                      </p>
                    </div>
                    
                    <div className="space-y-4">
                      <p className="text-muted-foreground leading-relaxed">
                        Dr. Isaiah Macwealth is a Christian leader, author, and philanthropist with a strong passion for spiritual growth and societal transformation. As Senior Pastor of Gospel Pillars International Churches and founder of OneSound Revival Fellowship, he has established platforms that promote faith, leadership development, education, and humanitarian outreach. His work spans media, ministerial training, and charitable initiatives that impact lives across nations.
                      </p>
                      
                      <p className="text-muted-foreground leading-relaxed">
                        In 2023, he launched the Ark of Light for All Nations in Lagos, a vibrant worship and community hub that also houses the Ark Food Bank, providing practical support through food and emergency assistance programs. His leadership reflects a commitment not only to spiritual teaching but also to community empowerment and structured development.
                      </p>
                      
                      <p className="text-muted-foreground leading-relaxed">
                        As Founder of the U-Genius Club, Dr. Macwealth brings this same vision into the academic space, championing excellence, character formation, and intellectual growth. His passion is to raise disciplined, purpose-driven young leaders equipped to make meaningful contributions to their communities and the world.
                      </p>
                    </div>
                    
                   
                  </div>
                </div>
                
                <div className="order-1 lg:order-2">
                  <div className="relative">
                    <div className="aspect-square rounded-2xl overflow-hidden bg-gradient-to-br from-gold/20 to-coral/20">
                      <img 
                        src={prophet}
                        alt="Dr. Isaiah Macwealth - Founder of U-Genius"
                        className="w-full h-full object-cover"
                      />
                    </div>
                    
                    {/* Decorative elements */}
                    <div className="absolute -top-4 -right-4 w-24 h-24 bg-gold/10 rounded-full blur-xl" />
                    <div className="absolute -bottom-4 -left-4 w-32 h-32 bg-coral/10 rounded-full blur-xl" />
                  </div>
                </div>
              </div>
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
                Our vision is to build a global community of purpose-driven students who excel academically, lead boldly, and create lasting impact. We see a generation that is confident in their abilities, disciplined in their habits, and intentional about their growth. Through learning, leadership, and service, we aim to shape individuals who influence their campuses, communities, and the world for good.
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
