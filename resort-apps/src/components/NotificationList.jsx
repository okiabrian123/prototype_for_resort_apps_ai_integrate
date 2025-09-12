import React from 'react';

const NotificationList = ({ onClose }) => {
  const notifications = [
    {
      id: 1,
      title: 'Booking Confirmed',
      message: 'Your booking at The Serenity Resort has been confirmed.',
      time: '2 hours ago',
      unread: true
    },
    {
      id: 2,
      title: 'Special Offer',
      message: 'Get 20% off on your next booking. Limited time offer!',
      time: '1 day ago',
      unread: true
    },
    {
      id: 3,
      title: 'Payment Reminder',
      message: 'Your payment for The Oasis Retreat is due in 3 days.',
      time: '2 days ago',
      unread: false
    },
    {
      id: 4,
      title: 'Resort Update',
      message: 'The pool at Phuket Paradise Villas will be closed for maintenance next week.',
      time: '3 days ago',
      unread: false
    }
  ];

  return (
    <div className="absolute inset-0 z-50 flex items-start justify-end pt-20 pr-4">
      <div className="bg-white rounded-xl shadow-lg w-full max-w-sm border border-gray-200">
        <div className="p-4 border-b border-gray-200 flex items-center justify-between">
          <h2 className="text-lg font-bold text-gray-900">Notifications</h2>
          <button 
            onClick={onClose}
            className="text-gray-500 hover:text-gray-700"
          >
            <span className="material-symbols-outlined">close</span>
          </button>
        </div>
        <div className="max-h-96 overflow-y-auto">
          {notifications.length === 0 ? (
            <div className="p-8 text-center">
              <span className="material-symbols-outlined text-4xl text-gray-300 mb-2">notifications</span>
              <p className="text-gray-500">No notifications yet</p>
            </div>
          ) : (
            <ul>
              {notifications.map((notification) => (
                <li 
                  key={notification.id} 
                  className={`p-4 border-b border-gray-100 hover:bg-gray-50 cursor-pointer ${notification.unread ? 'bg-blue-50' : ''}`}
                >
                  <div className="flex justify-between">
                    <h3 className="font-semibold text-gray-900">{notification.title}</h3>
                    <span className="text-xs text-gray-500">{notification.time}</span>
                  </div>
                  <p className="text-gray-600 text-sm mt-1">{notification.message}</p>
                </li>
              ))}
            </ul>
          )}
        </div>
        <div className="p-3 border-t border-gray-200 text-center">
          <button className="text-sm font-medium text-blue-600 hover:text-blue-800">
            Mark all as read
          </button>
        </div>
      </div>
      <div 
        className="fixed inset-0 bg-black/50 z-[-1]" 
        onClick={onClose}
      ></div>
    </div>
  );
};

export default NotificationList;