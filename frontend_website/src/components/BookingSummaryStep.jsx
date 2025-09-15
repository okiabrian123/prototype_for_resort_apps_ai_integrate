import React from 'react';

const BookingSummaryStep = ({ bookingData, onConfirm, onCancel }) => {
  const { date, guests, houseType } = bookingData;

  return (
    <div className="rounded-2xl px-4 py-3 max-w-[100%] bg-gray-100 text-gray-900 rounded-bl-lg">
      Please confirm or cancel this booking.
      <div className="mt-4 border border-gray-200 rounded-lg p-4 bg-white">
        <h3 className="text-lg font-semibold text-gray-800 mb-3">Booking Summary</h3>
        <div className="space-y-2 mb-4">
          <div className="flex justify-between">
            <span className="text-gray-600">Date:</span>
            <span className="font-medium">{date}</span>
          </div>
          <div className="flex justify-between">
            <span className="text-gray-600">Guests:</span>
            <span className="font-medium">{guests} people</span>
          </div>
          <div className="flex justify-between">
            <span className="text-gray-600">House Type:</span>
            <span className="font-medium">{houseType}</span>
          </div>
        </div>
        <div className="flex space-x-3">
          <button
            onClick={onCancel}
            className="flex-1 bg-gray-200 hover:bg-gray-300 text-gray-800 font-medium py-2 px-3 rounded-lg transition duration-200 text-sm"
          >
            Cancel
          </button>
          <button
            onClick={onConfirm}
            className="flex-1 bg-green-500 hover:bg-green-600 text-white font-medium py-2 px-3 rounded-lg transition duration-200 text-sm"
          >
            Confirm
          </button>
        </div>
      </div>
    </div>
  );
};

export default BookingSummaryStep;