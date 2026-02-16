import { useEffect, useState } from 'react';

export default function LoadingSpinner() {
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const timer = setTimeout(() => {
      setIsLoading(false);
    }, 2000); // Show loading for 2 seconds

    return () => clearTimeout(timer);
  }, []);

  if (!isLoading) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-gradient-to-br from-gold-50 to-white">
      <div className="text-center">
        <div className="relative w-32 h-32 mx-auto mb-8">
          {/* Logo SVG */}
          <svg
            className="w-full h-full animate-pulse"
            viewBox="0 0 100 100"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            {/* U-Genius Logo */}
            <circle cx="50" cy="50" r="45" className="stroke-gold-600 stroke-2 fill-none" />
            <path
              d="M30 35 L30 65 L40 65 L40 45 L45 45 L45 65 L55 65 L55 45 L60 45 L60 65 L70 65 L70 35 L60 35 L60 40 L55 40 L55 35 L45 35 L45 40 L40 40 L40 35 Z"
              className="fill-gold-600"
            />
            <circle cx="50" cy="25" r="3" className="fill-gold-600" />
            <circle cx="25" cy="50" r="3" className="fill-gold-600" />
            <circle cx="75" cy="50" r="3" className="fill-gold-600" />
            <circle cx="50" cy="75" r="3" className="fill-gold-600" />
          </svg>
          
          {/* Rotating ring */}
          <div className="absolute inset-0 border-4 border-gold-200 border-t-gold-600 rounded-full animate-spin"></div>
        </div>
        
        <h1 className="text-2xl font-bold text-gray-900 mb-2">U-Genius</h1>
        <p className="text-gray-600 animate-pulse">Loading your academic journey...</p>
        
        {/* Progress dots */}
        <div className="flex justify-center mt-4 space-x-2">
          <div className="w-2 h-2 bg-gold-600 rounded-full animate-bounce" style={{ animationDelay: '0ms' }}></div>
          <div className="w-2 h-2 bg-gold-600 rounded-full animate-bounce" style={{ animationDelay: '150ms' }}></div>
          <div className="w-2 h-2 bg-gold-600 rounded-full animate-bounce" style={{ animationDelay: '300ms' }}></div>
        </div>
      </div>
    </div>
  );
}
