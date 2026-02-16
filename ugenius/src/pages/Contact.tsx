import { useState } from "react";
import Navbar from "@/components/layout/Navbar";
import Footer from "@/components/layout/Footer";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { useToast } from "@/hooks/use-toast";
import { Mail, MapPin, Phone, Send, Facebook, Twitter, Instagram, Linkedin } from "lucide-react";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";

const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080/api/v1";

const faqs = [
  {
    question: "How do I become a member of U-Genius?",
    answer: "Simply fill out the registration form on our Join page. Once submitted, our team will review your application and reach out within 48 hours with next steps.",
  },
  {
    question: "Is there a membership fee?",
    answer: "Basic membership is free! However, we offer premium tiers with additional benefits such as one-on-one mentorship and exclusive resources for a nominal fee.",
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
  const { toast } = useToast();
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    subject: "",
    message: "",
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setFormData((prev) => ({
      ...prev,
      [e.target.name]: e.target.value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSubmitting(true);

    try {
      // Send to your Go backend
      const response = await fetch(`${API_URL}/contact`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          name: formData.name,
          email: formData.email,
          subject: formData.subject,
          message: formData.message,
        }),
      });

      if (response.ok) {
        toast({
          title: "Message Sent!",
          description: "We'll get back to you as soon as possible.",
        });
        setFormData({ name: "", email: "", subject: "", message: "" });
      } else {
        throw new Error("Failed to send message");
      }
    } catch {
      toast({
        title: "Error",
        description: "Failed to send message. Please try again later.",
        variant: "destructive",
      });
    } finally {
      setIsSubmitting(false);
    }
  };

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
                Get in Touch
              </span>
              <h1 className="font-display text-5xl md:text-6xl font-bold text-foreground mt-4 mb-6">
                Contact Us
              </h1>
              <p className="text-lg text-muted-foreground leading-relaxed">
                Have questions about U-Genius? We'd love to hear from you. 
                Send us a message and we'll respond as soon as possible.
              </p>
            </div>
          </div>
        </section>

        {/* Contact Section */}
        <section className="py-24 bg-background">
          <div className="container mx-auto px-4">
            <div className="max-w-6xl mx-auto">
              <div className="grid grid-cols-1 lg:grid-cols-2 gap-16">
                {/* Contact Info */}
                <div>
                  <h2 className="font-display text-3xl font-bold text-foreground mb-8">
                    Reach Out to Us
                  </h2>

                  <div className="space-y-6 mb-12">
                    <div className="flex items-start gap-4">
                      <div className="w-12 h-12 rounded-xl bg-gold/10 flex items-center justify-center shrink-0">
                        <MapPin className="w-6 h-6 text-gold" />
                      </div>
                      <div>
                        <h3 className="font-semibold text-foreground mb-1">Location</h3>
                        <p className="text-muted-foreground">
                          Student Center, Main Campus<br />
                          University Avenue
                        </p>
                      </div>
                    </div>

                    <div className="flex items-start gap-4">
                      <div className="w-12 h-12 rounded-xl bg-gold/10 flex items-center justify-center shrink-0">
                        <Mail className="w-6 h-6 text-gold" />
                      </div>
                      <div>
                        <h3 className="font-semibold text-foreground mb-1">Email</h3>
                        <a href="mailto:info@ugenius.club" className="text-muted-foreground hover:text-gold transition-colors">
                          info@ugenius.club
                        </a>
                      </div>
                    </div>

                    <div className="flex items-start gap-4">
                      <div className="w-12 h-12 rounded-xl bg-gold/10 flex items-center justify-center shrink-0">
                        <Phone className="w-6 h-6 text-gold" />
                      </div>
                      <div>
                        <h3 className="font-semibold text-foreground mb-1">Phone</h3>
                        <a href="tel:+1234567890" className="text-muted-foreground hover:text-gold transition-colors">
                          +123 456 7890
                        </a>
                      </div>
                    </div>
                  </div>

                  {/* Social Links */}
                  <div>
                    <h3 className="font-semibold text-foreground mb-4">Follow Us</h3>
                    <div className="flex gap-3">
                      {[Facebook, Twitter, Instagram, Linkedin].map((Icon, i) => (
                        <a
                          key={i}
                          href="#"
                          className="w-12 h-12 rounded-xl bg-gold/10 flex items-center justify-center text-gold hover:bg-gold hover:text-primary transition-all duration-300"
                        >
                          <Icon size={20} />
                        </a>
                      ))}
                    </div>
                  </div>
                </div>

                {/* Contact Form */}
                <div className="bg-card rounded-3xl p-8 md:p-10 shadow-soft">
                  <h2 className="font-display text-2xl font-semibold text-foreground mb-6">
                    Send a Message
                  </h2>

                  <form onSubmit={handleSubmit} className="space-y-6">
                    <div className="space-y-2">
                      <Label htmlFor="name">Your Name</Label>
                      <Input
                        id="name"
                        name="name"
                        value={formData.name}
                        onChange={handleChange}
                        required
                        placeholder="John Doe"
                        className="bg-background"
                      />
                    </div>

                    <div className="space-y-2">
                      <Label htmlFor="email">Email Address</Label>
                      <Input
                        id="email"
                        name="email"
                        type="email"
                        value={formData.email}
                        onChange={handleChange}
                        required
                        placeholder="john@example.com"
                        className="bg-background"
                      />
                    </div>

                    <div className="space-y-2">
                      <Label htmlFor="subject">Subject</Label>
                      <Input
                        id="subject"
                        name="subject"
                        value={formData.subject}
                        onChange={handleChange}
                        required
                        placeholder="How can we help?"
                        className="bg-background"
                      />
                    </div>

                    <div className="space-y-2">
                      <Label htmlFor="message">Message</Label>
                      <Textarea
                        id="message"
                        name="message"
                        value={formData.message}
                        onChange={handleChange}
                        required
                        placeholder="Tell us more about your inquiry..."
                        rows={5}
                        className="bg-background resize-none"
                      />
                    </div>

                    <Button
                      type="submit"
                      variant="hero"
                      size="xl"
                      className="w-full"
                      disabled={isSubmitting}
                    >
                      {isSubmitting ? (
                        <span className="flex items-center gap-2">
                          <div className="w-5 h-5 border-2 border-primary border-t-transparent rounded-full animate-spin" />
                          Sending...
                        </span>
                      ) : (
                        <span className="flex items-center gap-2">
                          <Send className="w-5 h-5" />
                          Send Message
                        </span>
                      )}
                    </Button>
                  </form>
                </div>
              </div>
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
