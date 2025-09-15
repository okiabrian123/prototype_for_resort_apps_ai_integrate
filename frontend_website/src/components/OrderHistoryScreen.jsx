import React from 'react';

const OrderHistoryScreen = ({ navigateTo }) => {
  const orders = [
    {
      id: 1,
      name: 'The Serenity Resort',
      date: 'Oct 15-18, 2023',
      image: 'https://lh3.googleusercontent.com/aida-public/AB6AXuBq951Cbj4UT8esgvgWx2stU2vsfyaqcSaLJ-UqSFZTIHS7UhbJP_d737b7_tnW2IrwgoHR85UKy1ESBUMkh_ZQ-7g0Rpf1dSfKQLbiWFoPEg_pMlRf1J8NVRx6nu5JYqian-hxhZmLDHeY1qJNGtc4Xd4bRmMlrPjTjpOEbDAOGoWfXu0wcecLUxcKvqUzJ-bSLjMITRVpTiMNFhYKg6qJPX2YwoqF2BdvnVf2cuPhYrFi-lcKahFugSENQwdRByi4C4XO2EYRnt6C',
      status: 'Pending',
      statusClass: 'bg-yellow-500/20 text-yellow-400'
    },
    {
      id: 2,
      name: 'The Oasis Retreat',
      date: 'Sep 20-25, 2023',
      image: 'https://lh3.googleusercontent.com/aida-public/AB6AXuD5lu_gIF8d2DDY1MPw5R--MBri4NycTy8OCLstwqdYahl2qCoXYwT4JMd1Y5nR9T0gaI2hypncFfzBzuWidMiXYdI2IRIXF5CL2SOMj6VTjGdoM2wSeDrNldq136bp9Gnu_zlvwCDDqecvAHzyuInRcXccIdRuMnvGrJwlqP58OnvYTuZ24smQYe07eOj9BY5BMg_JV-0tCE5H4pMPHONpenk45M7wp33-nizH2UhT2qAHnHFaASjI6kK4bGkKN6QPT2P9OSx3cVQ6',
      status: 'Confirmed',
      statusClass: 'bg-green-500/20 text-green-400'
    },
    {
      id: 3,
      name: 'Phuket Paradise Villas',
      date: 'Jul 5-10, 2023',
      image: 'https://lh3.googleusercontent.com/aida-public/AB6AXuA-T1pbUFSkOKLfXnN3umgXTK8w1DZDwX45LmQH9YflmSQ5lzSdbxkEen5YiINte4LjZojQuACRXSNqmWu8kEdTd3U3n42z5cEhzr2WxL_figz6WQ6cElYPxGLVIQYFk4Azn13_9o2jxbIUYK4yAuN3qqhjwovE4LSQ4h9cblxLWQpWfunCD2yrjx2Ge8erk8GL3yM43-bMOY5wUsmDFKv8vXhp9Kdcxr8p9uvj_doIf7o70jD1O8cqpsvN5ToX9rxvPOVooXkk-FEq',
      status: 'Completed',
      statusClass: 'bg-blue-500/20 text-blue-400'
    },
    {
      id: 4,
      name: 'Krabi Beachfront Bungalow',
      date: 'May 1-3, 2023',
      image: 'https://lh3.googleusercontent.com/aida-public/AB6AXuCgWmqeKywdw_vZDF9zRzItrjhgFrleq3nrt7NFiAyClFZ4o54eXJsChc0vbAEQkGRpUQ5G7UFk-f5SrJeLTx7qlAA_9L9nNjLnCU28zg_ssqHPCEp3PAQb2jkIN1SrV39_FDprhX7XdEFeBILeLK7uMK3hnc2z0Dwei8MV0bmlQF0tI_81i45OCYXssU6xQxNViD4LUMhx2S3nibAGO0V5vUdg94YJxiU476ow3Y4HjoDL9ljGxi820Fas4Hyx4-Z7oPJ4MXaAbQIR',
      status: 'Cancelled',
      statusClass: 'bg-red-500/20 text-red-400'
    }
  ];

  return (
    <div className="relative flex size-full min-h-screen flex-col bg-white justify-between group/design-root overflow-x-hidden text-zinc-900">
      <div className="flex-grow">
        <div className="flex items-center p-4 justify-between sticky top-0 bg-white z-10">
          <button 
            className="flex size-10 shrink-0 items-center justify-center rounded-full text-zinc-600 hover:bg-zinc-100 hover:text-zinc-900 transition-colors"
            onClick={() => navigateTo && navigateTo('home')}
          >
            <span className="material-symbols-outlined">arrow_back</span>
          </button>
          <h2 className="text-xl font-bold leading-tight tracking-tight text-center text-zinc-900">Booking History</h2>
          <div className="size-10"></div>
        </div>
        <div className="p-4 space-y-4">
          {orders.map((order) => (
            <div key={order.id} className="flex items-center gap-4 bg-zinc-50 p-4 rounded-lg">
              <div 
                className="bg-center bg-no-repeat aspect-square bg-cover rounded-md h-16 w-16" 
                style={{ backgroundImage: `url("${order.image}")` }}
              ></div>
              <div className="flex-1">
                <p className="text-lg font-semibold leading-normal text-zinc-900">{order.name}</p>
                <div className="flex items-center gap-2 mt-1">
                  <span className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${order.statusClass}`}>
                    {order.status}
                  </span>
                  <p className="text-zinc-600 text-sm">{order.date}</p>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default OrderHistoryScreen;