"use client";

import DeviceList from "@/components/DeviceList";
import React from "react";

export default function Home() {
  const [devices, setDevices] = React.useState<Device[]>([]);

  React.useEffect(() => {
    const fetchDevices = async () => {
      const response = await fetch("/api/devices");
      const data = await response.json();
      setDevices(data);
    };
    fetchDevices();
  }, []);

  return (
    <div>
      <main>
        <DeviceList devices={devices} type="card" />
      </main>
    </div>
  );
}
