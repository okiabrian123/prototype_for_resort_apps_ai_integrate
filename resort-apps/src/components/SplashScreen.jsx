import React, { useEffect } from 'react';

const SplashScreen = ({ onFinished }) => {
  useEffect(() => {
    const timer = setTimeout(() => {
      if (onFinished) {
        onFinished();
      }
    }, 2500);

    return () => clearTimeout(timer);
  }, [onFinished]);

  return (
    <div className="relative flex size-full min-h-screen flex-col items-center justify-center bg-[#F0FFF4] dark:bg-[#1A2E22] group/design-root overflow-hidden fade-in">
      <div className="flex flex-col items-center gap-4">
        <svg className="h-24 w-24 text-[#38e07b]" fill="none" height="96" viewBox="0 0 24 24" width="96" xmlns="http://www.w3.org/2000/svg">
          <path d="M12 2L1 9L4 22H20L23 9L12 2Z" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
          <path d="M12 2V22" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
          <path d="M23 9H1" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
          <path d="M4 22L12 15L20 22" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
          <path d="M10 12H14" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
        </svg>
        <h1 className="text-3xl font-bold text-[#1A2E22] dark:text-white">Thai Serenity</h1>
        <p className="text-lg text-gray-600 dark:text-gray-300">Your gateway to paradise</p>
      </div>
    </div>
  );
};

export default SplashScreen;