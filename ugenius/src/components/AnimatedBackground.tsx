import { motion } from 'framer-motion';

export default function AnimatedBackground() {
  const particles = [
    { id: 1, left: 10, top: 20, duration: 3, delay: 0 },
    { id: 2, left: 30, top: 60, duration: 4, delay: 0.5 },
    { id: 3, left: 50, top: 40, duration: 5, delay: 1 },
    { id: 4, left: 70, top: 80, duration: 3.5, delay: 1.5 },
    { id: 5, left: 90, top: 30, duration: 4.5, delay: 0.2 },
    { id: 6, left: 20, top: 90, duration: 6, delay: 0.8 },
    { id: 7, left: 40, top: 10, duration: 3.2, delay: 1.2 },
    { id: 8, left: 60, top: 70, duration: 5.5, delay: 0.3 },
    { id: 9, left: 80, top: 50, duration: 4.8, delay: 0.7 },
    { id: 10, left: 15, top: 35, duration: 3.7, delay: 1.1 },
    { id: 11, left: 35, top: 75, duration: 5.2, delay: 0.4 },
    { id: 12, left: 55, top: 25, duration: 4.3, delay: 0.9 },
    { id: 13, left: 75, top: 85, duration: 3.9, delay: 1.3 },
    { id: 14, left: 25, top: 55, duration: 5.8, delay: 0.1 },
    { id: 15, left: 45, top: 95, duration: 4.1, delay: 0.6 },
    { id: 16, left: 65, top: 15, duration: 3.4, delay: 1.4 },
    { id: 17, left: 85, top: 65, duration: 5.1, delay: 0.2 },
    { id: 18, left: 5, top: 45, duration: 4.6, delay: 0.8 },
    { id: 19, left: 95, top: 5, duration: 3.3, delay: 1.6 },
    { id: 20, left: 50, top: 50, duration: 5.9, delay: 0.5 },
  ];

  return (
    <div className="fixed inset-0 -z-10 overflow-hidden">
      {/* Animated gradient orbs */}
      <motion.div
        className="absolute top-0 left-0 w-96 h-96 bg-gold-200 rounded-full mix-blend-multiply filter blur-xl opacity-70"
        animate={{
          x: [0, 100, 0],
          y: [0, -100, 0],
        }}
        transition={{
          duration: 20,
          repeat: Infinity,
          repeatType: "reverse",
        }}
      />
      
      <motion.div
        className="absolute top-0 right-0 w-96 h-96 bg-purple-200 rounded-full mix-blend-multiply filter blur-xl opacity-70"
        animate={{
          x: [0, -100, 0],
          y: [0, 100, 0],
        }}
        transition={{
          duration: 25,
          repeat: Infinity,
          repeatType: "reverse",
        }}
      />
      
      <motion.div
        className="absolute bottom-0 left-1/2 w-96 h-96 bg-blue-200 rounded-full mix-blend-multiply filter blur-xl opacity-70"
        animate={{
          x: [-50, 50, -50],
          y: [0, -50, 0],
        }}
        transition={{
          duration: 30,
          repeat: Infinity,
          repeatType: "reverse",
        }}
      />

      {/* Floating particles */}
      {particles.map((particle) => (
        <motion.div
          key={particle.id}
          className="absolute w-2 h-2 bg-gold-400 rounded-full opacity-60"
          style={{
            left: `${particle.left}%`,
            top: `${particle.top}%`,
          }}
          animate={{
            y: [0, -100, 0],
            opacity: [0.6, 1, 0.6],
          }}
          transition={{
            duration: particle.duration,
            repeat: Infinity,
            repeatType: "reverse",
            delay: particle.delay,
          }}
        />
      ))}

      {/* Gradient overlay */}
      <div className="absolute inset-0 bg-gradient-to-br from-gold-50/30 via-white/50 to-purple-50/30" />
    </div>
  );
}
