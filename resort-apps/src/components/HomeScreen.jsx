import React, { useState } from 'react';
import NotificationList from './NotificationList';

const HomeScreen = ({ navigateTo }) => {
  const [showNotifications, setShowNotifications] = useState(false);

  return (
    <div className="relative flex size-full min-h-screen flex-col group/design-root overflow-x-hidden">
      {showNotifications && (
        <NotificationList onClose={() => setShowNotifications(false)} />
      )}
      
      <div 
        className="h-80 w-full bg-image"
        style={{ 
          backgroundImage: "url('https://images.unsplash.com/photo-1540541338287-41700207dee6?q=80&w=2940&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D')",
          backgroundSize: 'cover',
          backgroundPosition: 'center'
        }}
      >
        <div className="bg-black/20 h-full w-full p-6 flex flex-col">
          <div className="flex items-center justify-end">
            <button 
              className="relative flex items-center justify-center rounded-full h-12 w-12 text-white hover:bg-white/20 transition-colors"
              onClick={() => setShowNotifications(true)}
            >
              <span className="material-symbols-outlined text-3xl">
                notifications
              </span>
              <span className="absolute top-2 right-2 flex h-3 w-3">
                <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-red-400 opacity-75"></span>
                <span className="relative inline-flex rounded-full h-3 w-3 bg-red-500"></span>
              </span>
            </button>
          </div>
          <div className="text-white">
            <h1 className="text-3xl font-bold tracking-tight">Hi,</h1>
            <p className="text-white/90 text-2xl mt-1">Ready to book your <br/> next getaway?</p>
          </div>
        </div>
      </div>
      
      <main className="flex-1 bg-gray-50 rounded-t-3xl -mt-8 p-6">
        <div className="grid grid-cols-2 gap-4">
          <button 
            onClick={() => navigateTo && navigateTo('chat')}
            className="group flex flex-col items-center justify-center rounded-2xl bg-white p-[10px] aspect-square shadow-sm hover:shadow-lg transition-shadow duration-300"
          >
            <div className="flex items-center justify-center h-16 w-16 rounded-full bg-green-100 text-green-600 mb-4 transition-transform duration-300 group-hover:scale-110">
              <span className="material-symbols-outlined text-4xl">
                chat_bubble
              </span>
            </div>
            <h2 className="text-lg font-semibold text-gray-800 text-center">Chat Booking</h2>
            <p className="text-sm text-gray-500 text-center mt-1">Start a new booking</p>
          </button>
          
          <button 
            onClick={() => navigateTo && navigateTo('orders')}
            className="group flex flex-col items-center justify-center rounded-2xl bg-white p-[10px] aspect-square shadow-sm hover:shadow-lg transition-shadow duration-300"
          >
            <div className="flex items-center justify-center h-16 w-16 rounded-full bg-purple-100 text-purple-600 mb-4 transition-transform duration-300 group-hover:scale-110">
              <span className="material-symbols-outlined text-4xl">
                history
              </span>
            </div>
            <h2 className="text-lg font-semibold text-gray-800 text-center">Order History</h2>
            <p className="text-sm text-gray-500 text-center mt-1">View past bookings</p>
          </button>
        </div>
        
        <div className="mt-8">
          <h3 className="text-xl font-bold text-gray-800 mb-4">Promotions</h3>
          <div className="space-y-4">
            <div className="bg-white p-4 rounded-xl shadow-sm flex items-center gap-4">
              <img 
                alt="Resort Promotion" 
                className="w-24 h-24 object-cover rounded-lg" 
                src="https://lh3.googleusercontent.com/aida-public/AB6AXuBS4ql2j4PrEv_CaidmPRWDoiZblANj5J0O3Q22f2wcz_AVtmWXZQA_QJuxrjmmIJ4czvbNGgT1kiTmUQTmCbodgCzljqQdn_tCHBl2a6HsIM6hlFxgcJuMksFC7IdAj_HorVj2ypHs6Na0wcfudvEhyxeFLnJLI5x5EFED3MuIKNF3y69OWKxJdOWa2KG0IM4CxDcahxXUomNNZeysJTeR7sivRE1MQRQCw6U1eVXeU4jcO78D0DtVXjTT4G-Ras0FBrHZiJS2z083" 
              />
              <div>
                <h4 className="font-semibold text-gray-800">30% Off Your Next Stay</h4>
                <p className="text-sm text-gray-600 mt-1">Book before the end of the month to get a discount.</p>
                <button className="text-sm font-semibold text-primary-600 mt-2 hover:underline">Book now</button>
              </div>
            </div>
            
            <div className="bg-white p-4 rounded-xl shadow-sm flex items-center gap-4">
              <img 
                alt="Resort Promotion" 
                className="w-24 h-24 object-cover rounded-lg" 
                src="https://lh3.googleusercontent.com/aida-public/AB6AXuCzKKdMppUOj7P538YXim3Jql9fpQVcIRs0Vd_apIykJO5Baxt9RtwhZ7Q_lkK9RPXBZydvRdu2mNGt-EuUHG98D2g9KvRu2J-8je0DuBg1senHRzeAH0tNg5j6dgL6tdaNYGZVAgROzplb6Did2c1WvipCWIJ9CztldHoV6-MFE4ZSaP77VuAmsUGb6g1LieElexB4pkOb_ZwwmjO1BeNIDrx06iA4hlaBGyD3b0l35irwG28ri8ntu1hSVw0BMd1Bw184cWhcFW2I" 
              />
              <div>
                <h4 className="font-semibold text-gray-800">Free Spa Treatment</h4>
                <p className="text-sm text-gray-600 mt-1">Enjoy a complimentary spa session with any 3-night stay.</p>
                <button className="text-sm font-semibold text-primary-600 mt-2 hover:underline">Learn more</button>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
};

export default HomeScreen;