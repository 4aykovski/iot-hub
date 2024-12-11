"use client";

import DeviceList from "@/components/DeviceList";
import React from "react";

export default function Devices() {
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
      <main className="flex justify-center items-center">
        <DeviceList devices={devices} type="table" />
      </main>
    </div>
  );
}
