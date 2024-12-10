import DeviceList from "@/components/DeviceList";
import React from "react";

export default function Home() {
  const [devices, setDevices] = React.useState<Device[]>([]);
  const [data, setData] = React.useState([]);

  React.useEffect(() => {
    const fetchDevices = async () => {
      const response = await fetch("/api/v1/devices");
      const data = await response.json();
      setDevices(data);
    };
    fetchDevices();
  });

  return (
    <div>
      <main>
        <DeviceList devices={devices} type="card" />
      </main>
    </div>
  );
}
