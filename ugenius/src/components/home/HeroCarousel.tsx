import { useState, useEffect } from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { ArrowRight, ChevronLeft, ChevronRight } from 'lucide-react';
import { Button } from '@/components/ui/button';
import image1 from "@/assets/image1new.jpg"
import image2 from "@/assets/image2new.jpg"
import image3 from "@/assets/image4new.jpg"
import image4 from "@/assets/image3new.jpg"


const slides = [
  {
    id: 1,
    image: image1,
    title: 'Unlock Your Academic Genius Within',
    highlighted: 'Academic',
    subtitle: 'Genius Within',
    description: 'Discover personalized learning paths tailored to your unique academic journey.',
  },
  {
    id: 2,
    image: image2,
    title: 'Transform Your Study Habits',
    highlighted: 'Study',
    subtitle: 'Habits',
    description: 'Build effective study routines that lead to consistent academic success.',
  },
  {
    id: 3,
    image: image3,
    title: 'Join a Community of Excellence Achievers',
    highlighted: 'Excellence',
    subtitle: 'Achievers',
    description: 'Connect with motivated students and achieve your goals together.',
  },
  {
    id: 4,
    image: image4,
    title: 'Achieve Your Dream Goals',
    highlighted: 'Dream',
    subtitle: 'Goals',
    description: 'Turn your academic aspirations into reality with expert guidance.',
  },
];

export default function HeroCarousel() {
  const [currentSlide, setCurrentSlide] = useState(0);
  const [isAutoPlaying,] = useState(true);

  useEffect(() => {
    if (!isAutoPlaying) return;

    const interval = setInterval(() => {
      setCurrentSlide((prev) => (prev + 1) % slides.length);
    }, 6000);

    return () => clearInterval(interval);
  }, [isAutoPlaying]);

  const nextSlide = () => {
    setCurrentSlide((prev) => (prev + 1) % slides.length);
  };

  const prevSlide = () => {
    setCurrentSlide((prev) => (prev - 1 + slides.length) % slides.length);
  };

  const goToSlide = (index: number) => {
    setCurrentSlide(index);
  };

  return (
    <div className="relative min-h-screen flex items-center justify-center overflow-hidden pt-8">
      {/* Carousel Slides */}
      <AnimatePresence mode="wait">
        <motion.div
          key={currentSlide}
          initial={{ opacity: 0, x: 1000 }}
          animate={{ opacity: 1, x: 0 }}
          exit={{ opacity: 0, x: -1000 }}
          transition={{ ease: 'easeInOut', duration: 0.5 }}
          className="absolute inset-0"
        >
          {/* Background Image */}
          <div className="absolute inset-0">
            <img
              src={slides[currentSlide].image}
              alt={slides[currentSlide].title}
              className="w-full h-full object-cover"
              loading="eager"
              decoding="async"
            />
            <div className="absolute inset-0 bg-gradient-to-b from-black/40 via-black/60 to-black/80" />
          </div>

          {/* Content */}
          <div className="relative z-10 container mx-auto px-4 pt-16">
            <div className="grid grid-cols-1 lg:grid-cols-2 gap-12 items-center min-h-[calc(100vh-8rem)]">
              {/* Left Column - Text */}
              <div className="text-center lg:text-left lg:pl-10">
                {/* Heading */}
                <motion.h1
                  initial={{ opacity: 0, y: 30 }}
                  animate={{ opacity: 1, y: 0 }}
                  transition={{ delay: 0.2, duration: 0.8 }}
                  className="font-display text-5xl md:text-6xl lg:text-6xl font-bold text-white mb-4"
                 
                >
                  {slides[currentSlide].title}
                  <span className="text-transparent bg-clip-text bg-gradient-to-r from-gold-400 to-gold-600">
                    {slides[currentSlide].highlighted}
                  </span>
                </motion.h1>

                {/* Subheading */}
                <motion.p
                  initial={{ opacity: 0, y: 30 }}
                  animate={{ opacity: 1, y: 0 }}
                  transition={{ delay: 0.4, duration: 0.8 }}
                  className="text-lg md:text-xl text-gray-200 max-w-2xl mx-auto lg:mx-0 mb-10 leading-relaxed"
                >
                  {slides[currentSlide].description}
                </motion.p>

                {/* CTA Buttons */}
                <motion.div
                  initial={{ opacity: 0, y: 30 }}
                  animate={{ opacity: 1, y: 0 }}
                  transition={{ delay: 0.6, duration: 0.8 }}
                  className="flex flex-col sm:flex-row gap-4 justify-center lg:justify-start"
                >
                  <Button variant="hero" size="xl" className="group">
                    Join U-Genius Today
                    <ArrowRight className="w-5 h-5 group-hover:translate-x-1 transition-transform" />
                  </Button>
                 
                </motion.div>

                {/* Stats */}
              
              </div>

              {/* Right Column - Empty for spacing */}
              <div className="hidden lg:block" />
            </div>
          </div>
        </motion.div>
      </AnimatePresence>

      {/* Navigation Controls */}
      <div className="absolute z-20 flex items-center justify-between w-full px-4">
        <Button
          variant="ghost"
          size="icon"
          onClick={prevSlide}
          className="text-white hover:bg-white/20 backdrop-blur-sm"
        >
          <ChevronLeft className="w-6 h-6" />
        </Button>
        <Button
          variant="ghost"
          size="icon"
          onClick={nextSlide}
          className="text-white hover:bg-white/20 backdrop-blur-sm"
        >
          <ChevronRight className="w-6 h-6" />
        </Button>
      </div>

      {/* Slide Indicators */}
      <div className="absolute z-20 bottom-8 left-1/2 -translate-x-1/2 flex space-x-2">
        {slides.map((_, index) => (
          <button
            key={index}
            onClick={() => goToSlide(index)}
            className={`w-3 h-3 rounded-full transition-all duration-300 ${
              index === currentSlide
                ? 'bg-gold-400 w-8'
                : 'bg-white/50 hover:bg-white/70'
            }`}
          />
        ))}
      </div>

      {/* Auto-play Toggle */}
      

      {/* Scroll Indicator */}
      <motion.div
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        transition={{ delay: 1, duration: 1 }}
        className="absolute z-20 bottom-24 left-1/2 -translate-x-1/2"
      >
        <div className="w-6 h-10 rounded-full border-2 border-white/30 flex items-start justify-center p-1">
          <div className="w-1.5 h-3 rounded-full bg-gold-400 animate-pulse" />
        </div>
      </motion.div>
    </div>
  );
}
