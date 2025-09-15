import React, { useState } from 'react';
import SplashScreen from './components/SplashScreen';
import HomeScreen from './components/HomeScreen';
import BookingChatScreen from './components/BookingChatScreen';
import OrderHistoryScreen from './components/OrderHistoryScreen';

const App = () => {
  const [currentScreen, setCurrentScreen] = useState('splash'); // splash, home, chat, orders

  const handleSplashFinished = () => {
    setCurrentScreen('home');
  };

  const navigateTo = (screen) => {
    setCurrentScreen(screen);
  };

  return (
    <div className=" bg-gray-100 mobile-apps">
      {currentScreen === 'splash' && (
        <SplashScreen onFinished={handleSplashFinished} />
      )}
      
      {currentScreen === 'home' && (
        <HomeScreen navigateTo={navigateTo} />
      )}
      
      {currentScreen === 'chat' && <BookingChatScreen navigateTo={navigateTo} />}
      
      {currentScreen === 'orders' && <OrderHistoryScreen navigateTo={navigateTo} />}
    </div>
  );
};

export default App;