import Navbar from "@/components/layout/Navbar";
import Footer from "@/components/layout/Footer";
import HeroSection from "@/components/home/HeroSection";
import MissionSection from "@/components/home/MissionSection";
import BenefitsPreview from "@/components/home/BenefitsPreview";
// import EventsPreview from "@/components/home/EventsPreview";
import TestimonialsPreview from "@/components/home/TestimonialsPreview";
import CTASection from "@/components/home/CTASection";
import AnimatedBackground from "@/components/AnimatedBackground";

const Index = () => {
  return (
    <div className="min-h-screen relative">
      <AnimatedBackground />
      <Navbar />
      <main>
        <HeroSection />
        <MissionSection />
        <BenefitsPreview />
        {/* <EventsPreview /> */}
        <TestimonialsPreview />
        <CTASection />
      </main>
      <Footer />
    </div>
  );
};

export default Index;
