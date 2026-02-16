import { useState, useEffect } from 'react';
import { motion, AnimatePresence } from 'framer-motion';

export default function EntryLoading() {
  const hasVisited = localStorage.getItem('ugenius-has-visited');
  const [showLoading, setShowLoading] = useState(!hasVisited);
  const [showContent, setShowContent] = useState(!!hasVisited);

  useEffect(() => {
    if (!hasVisited) {
      // Mark as visited
      localStorage.setItem('ugenius-has-visited', 'true');
      
      // Hide loading after 3.5 seconds and show content
      setTimeout(() => {
        setShowLoading(false);
        setShowContent(true);
      }, 3500);
    }
  }, [hasVisited]);

  if (!showContent) {
    return (
      <AnimatePresence mode="wait">
        {showLoading && (
          <motion.div
            key="entry-loading"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.5 }}
            className="fixed inset-0 z-50 flex items-center justify-center bg-gradient-to-br from-gold-50 via-white to-purple-50"
          >
            <div className="text-center">
              {/* Logo Animation */}
              <motion.div
                initial={{ scale: 0, rotate: -180 }}
                animate={{ scale: 1, rotate: 0 }}
                transition={{ 
                  duration: 1.2, 
                  ease: "easeOut",
                  delay: 0.2 
                }}
                className="relative w-32 h-32 mx-auto mb-8"
              >
                {/* Logo SVG */}
                <svg
                  className="w-full h-full"
                  viewBox="0 0 100 100"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  {/* Outer ring */}
                  <motion.circle
                    cx="50"
                    cy="50"
                    r="45"
                    className="stroke-gold-600"
                    strokeWidth="2"
                    fill="none"
                    initial={{ pathLength: 0 }}
                    animate={{ pathLength: 1 }}
                    transition={{ duration: 1.5, delay: 0.5 }}
                  />
                  
                  {/* U-Genius Letter */}
                  <motion.path
                    d="M30 35 L30 65 L40 65 L40 45 L45 45 L45 65 L55 65 L55 45 L60 45 L60 65 L70 65 L70 35 L60 35 L60 40 L55 40 L55 35 L45 35 L45 40 L40 40 L40 35 Z"
                    className="fill-gold-600"
                    initial={{ pathLength: 0, opacity: 0 }}
                    animate={{ pathLength: 1, opacity: 1 }}
                    transition={{ duration: 1, delay: 1 }}
                  />
                  
                  {/* Decorative dots */}
                  {[0, 90, 180, 270].map((angle, i) => (
                    <motion.circle
                      key={i}
                      cx={50 + 25 * Math.cos((angle * Math.PI) / 180)}
                      cy={50 + 25 * Math.sin((angle * Math.PI) / 180)}
                      r="3"
                      className="fill-gold-600"
                      initial={{ scale: 0 }}
                      animate={{ scale: 1 }}
                      transition={{ duration: 0.3, delay: 1.2 + i * 0.1 }}
                    />
                  ))}
                </svg>
                
                {/* Rotating ring */}
                <motion.div
                  className="absolute inset-0 border-4 border-gold-200 border-t-gold-600 rounded-full"
                  animate={{ rotate: 360 }}
                  transition={{ duration: 2, repeat: Infinity, ease: "linear" }}
                />
              </motion.div>
              
              {/* Text Animation */}
              <motion.h1
                initial={{ opacity: 0, y: 20 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ duration: 0.8, delay: 1.5 }}
                className="text-4xl font-bold text-gray-900 mb-4"
              >
                U-Genius
              </motion.h1>
              
              <motion.p
                initial={{ opacity: 0, y: 20 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ duration: 0.8, delay: 1.8 }}
                className="text-lg text-gray-600 mb-8"
              >
                Unlocking Academic Excellence...
              </motion.p>
              
              {/* Progress dots */}
              <motion.div
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                transition={{ duration: 0.5, delay: 2 }}
                className="flex justify-center space-x-2"
              >
                {[0, 1, 2].map((i) => (
                  <motion.div
                    key={i}
                    className="w-3 h-3 bg-gold-600 rounded-full"
                    animate={{ scale: [1, 1.2, 1] }}
                    transition={{
                      duration: 1,
                      repeat: Infinity,
                      delay: i * 0.2,
                      repeatDelay: 0.5,
                    }}
                  />
                ))}
              </motion.div>
            </div>
          </motion.div>
        )}
      </AnimatePresence>
    );
  }

  // Show nothing - content will be rendered by the main app
  return null;
}
