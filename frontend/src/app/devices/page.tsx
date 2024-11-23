import DeviceList from "@/components/DeviceList";

export default function Devices() {
  const devices = [
    { name: "Device 1", data: "Data 1", ID: "1" },
    { name: "Device 2", data: "Data 2", ID: "2" },
    { name: "Device 3", data: "Data 3", ID: "3" },
    { name: "Device 4", data: "Data 4", ID: "4" },
    { name: "Device 5", data: "Data 5", ID: "5" },
    { name: "Device 6", data: "Data 6", ID: "6" },
  ];

  return (
    <div>
      <main className="flex justify-center items-center">
        <DeviceList devices={devices} type="table" />
      </main>
    </div>
  );
}
