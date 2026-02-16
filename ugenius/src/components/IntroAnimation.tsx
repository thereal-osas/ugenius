import { useState, useEffect } from 'react';
import { motion } from 'framer-motion';

const IntroAnimation = ({ onComplete }: { onComplete: () => void }) => {
  const [isVisible, setIsVisible] = useState(true);

  useEffect(() => {
    const timer = setTimeout(() => {
      setIsVisible(false);
      setTimeout(onComplete, 300); // Allow fade out animation to complete
    }, 3000); // 3 seconds

    return () => clearTimeout(timer);
  }, [onComplete]);

  if (!isVisible) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black">
      <motion.div
        initial={{ opacity: 0, scale: 0.5 }}
        animate={{ opacity: 1, scale: 1 }}
        exit={{ opacity: 0, scale: 0.8 }}
        transition={{ duration: 0.5, ease: "easeInOut" }}
        className="relative w-full h-full max-w-4xl max-h-screen"
      >
        {/* Video Background - Replace with your video file */}
        <video
          autoPlay
          muted
          playsInline
          className="w-full h-full object-cover"
          onEnded={() => {
            setIsVisible(false);
            setTimeout(onComplete, 300);
          }}
        >
          <source src="/assets/your_video_name.mp4" type="video/mp4" />
          {/* Fallback for browsers that don't support video */}
          Your browser does not support the video tag.
        </video>

        {/* Optional: Overlay text if needed */}
        <motion.div
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ duration: 0.5, delay: 1 }}
          className="absolute bottom-8 left-0 right-0 text-center"
        >
          <motion.p
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            transition={{ duration: 0.5, delay: 1.5 }}
            className="text-white/80 text-lg"
          >
            Loading...
          </motion.p>
        </motion.div>
      </motion.div>
    </div>
  );
};

export default IntroAnimation;
