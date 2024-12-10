"use client";

import TemperatureChart from "@/components/TemperatureChart";
import { useParams } from "next/navigation";
import React from "react";

const DevicePage: React.FC = () => {
  const [device, setDevice] = React.useState<Device | null>(null);
  const params = useParams<{ id: string }>();
  const id = params.id;

  React.useEffect(() => {
    const fetchDevice = async () => {
      const response = await fetch(`/api/v1/devices/${id}`);
      const data = await response.json();
      setDevice(data);
    };
    fetchDevice();
  }, [id]);

  return (
    <div className="mb-10 flex flex-col">
      <h4 className="mx-14 scroll-m-20 text-xl font-semibold tracking-tight">
        Device {id} - {device?.name}
      </h4>

      <div className="m-10 flex justify-center">
        <TemperatureChart id={id} className="w-full" />
      </div>
    </div>
  );
};

export default DevicePage;
