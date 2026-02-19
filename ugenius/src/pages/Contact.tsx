import Navbar from "@/components/layout/Navbar";
import Footer from "@/components/layout/Footer";
import { Mail, Phone } from "lucide-react";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";

const faqs = [
  {
    question: "How do I become a member of U-Genius?",
    answer: "Simply fill out the registration form on our Join page. Once submitted, our team will review your application and reach out within 48 hours with next steps.",
  },
  {
    question: "Is there a membership fee?",
    answer: "Membership is free! We offer premium benefits such as tutorials, one-on-one mentorship, skill-building initiatives, and exclusive resources for a no cost.",
  },
  {
    question: "What is Auto-Bio Jacking?",
    answer: "Auto-Bio Jacking is a transformative learning methodology developed by Dr. Isaiah Macwealth. It combines self-discovery, strategic thinking, and proven study techniques to help students achieve academic excellence.",
  },
  {
    question: "Can I join if I'm not a first-year student?",
    answer: "Absolutely! U-Genius welcomes students at all levels - from freshmen to postgraduates. It's never too late to transform your academic journey.",
  },
  {
    question: "How do the mentorship programs work?",
    answer: "Our mentorship program pairs you with successful graduates or senior students who've achieved academic excellence. You'll have regular meetings, study planning sessions, and guidance throughout your academic journey.",
  },
];

const Contact = () => {
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
            <div className="max-w-4xl mx-auto text-center">
              <span className="text-gold font-medium text-sm uppercase tracking-wider">
                Get in Touch
              </span>
              <h1 className="font-display text-5xl md:text-6xl font-bold text-foreground mt-4 mb-6">
                Contact Us
              </h1>
              <p className="text-lg text-muted-foreground leading-relaxed">
                Have questions about U-Genius? We'd love to hear from you. 
                Reach out to us through any of the channels below and we'll respond as soon as possible.
              </p>
            </div>
          </div>
        </section>

        {/* Contact Information Section */}
        <section className="py-24 bg-background">
          <div className="container mx-auto px-4">
            <div className="max-w-4xl mx-auto">
              <div className="text-center mb-16">
                <h2 className="font-display text-4xl font-bold text-foreground mb-6">
                  Reach Out to Us
                </h2>
                <p className="text-lg text-muted-foreground max-w-2xl mx-auto">
                  We're here to help you on your academic journey. Connect with us through any of these channels.
                </p>
              </div>

              {/* Contact Cards Grid */}
              <div className="grid grid-cols-1 md:grid-cols-2 gap-8 mb-16">
                {/* Location Card */}
                {/* <div className="text-center group">
                  <div className="w-20 h-20 rounded-2xl bg-gold/10 flex items-center justify-center mx-auto mb-6 group-hover:bg-gold/20 transition-colors">
                    <MapPin className="w-10 h-10 text-gold" />
                  </div>
                  <h3 className="font-semibold text-foreground text-xl mb-3">Location</h3>
                  <p className="text-muted-foreground leading-relaxed">
                    Student Center, Main Campus<br />
                    University Avenue
                  </p>
                </div> */}

                {/* Email Card */}
                <div className="text-center group">
                  <div className="w-20 h-20 rounded-2xl bg-gold/10 flex items-center justify-center mx-auto mb-6 group-hover:bg-gold/20 transition-colors">
                    <Mail className="w-10 h-10 text-gold" />
                  </div>
                  <h3 className="font-semibold text-foreground text-xl mb-3">Email</h3>
                  <a 
                    href="mailto:ugeniusclublagos@gmail.com" 
                    className="text-muted-foreground hover:text-gold transition-colors text-lg leading-relaxed"
                  >
                    ugeniusclublagos@gmail.com
                  </a>
                </div>

                {/* Phone Card */}
                <div className="text-center group">
                  <div className="w-20 h-20 rounded-2xl bg-gold/10 flex items-center justify-center mx-auto mb-6 group-hover:bg-gold/20 transition-colors">
                    <Phone className="w-10 h-10 text-gold" />
                  </div>
                  <h3 className="font-semibold text-foreground text-xl mb-3">Phone</h3>
                  <a 
                    href="tel:+2347062557356" 
                    className="text-muted-foreground hover:text-gold transition-colors text-lg leading-relaxed"
                  >
                    +234 706 255 7356
                  </a>
                </div>
              </div>

              {/* Social Media Section */}
              {/* <div className="text-center">
                <h3 className="font-semibold text-foreground text-xl mb-6">Follow Us</h3>
                <p className="text-muted-foreground mb-8 max-w-2xl mx-auto">
                  Stay connected with our community for updates, events, and inspiring content.
                </p>
                <div className="flex justify-center gap-4">
                  {[Facebook, Twitter, Instagram, Linkedin].map((Icon, i) => (
                    <a
                      key={i}
                      href="#"
                      className="w-14 h-14 rounded-2xl bg-gold/10 flex items-center justify-center text-gold hover:bg-gold hover:text-primary transition-all duration-300"
                    >
                      <Icon size={24} />
                    </a>
                  ))}
                </div>
              </div> */}
            </div>
          </div>
        </section>

        {/* FAQ Section */}
        <section className="py-24 bg-cream-dark">
          <div className="container mx-auto px-4">
            <div className="max-w-3xl mx-auto">
              <div className="text-center mb-12">
                <span className="text-gold font-medium text-sm uppercase tracking-wider">
                  FAQ
                </span>
                <h2 className="font-display text-4xl font-bold text-foreground mt-4">
                  Frequently Asked Questions
                </h2>
              </div>

              <Accordion type="single" collapsible className="space-y-4">
                {faqs.map((faq, index) => (
                  <AccordionItem
                    key={index}
                    value={`item-${index}`}
                    className="bg-card rounded-xl px-6 shadow-soft border-none"
                  >
                    <AccordionTrigger className="text-left font-semibold text-foreground hover:text-gold hover:no-underline py-5">
                      {faq.question}
                    </AccordionTrigger>
                    <AccordionContent className="text-muted-foreground pb-5">
                      {faq.answer}
                    </AccordionContent>
                  </AccordionItem>
                ))}
              </Accordion>
            </div>
          </div>
        </section>
      </main>
      <Footer />
    </div>
  );
};

export default Contact;
