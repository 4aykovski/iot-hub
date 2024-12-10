"use client";

import TemperatureChart from "@/components/TemperatureChart";
import { useParams } from "next/navigation";
import React from "react";
import UpdateDeviceBtn from "./UpdateDeviceBtn";

const DevicePage: React.FC = () => {
  const [device, setDevice] = React.useState<Device | null>(null);
  const params = useParams<{ id: string }>();
  const id = params.id;

  React.useEffect(() => {
    const fetchDevice = async () => {
      const response = await fetch(`/api/devices/${id}`);
      const data = await response.json();
      setDevice(data);
    };
    fetchDevice();
  }, [id]);

  return (
    <div className="mb-10 flex flex-col">
      <div className="flex justify-between mx-14">
        <h4 className=" scroll-m-20 text-xl font-semibold tracking-tight">
          Device {id} - {device?.name} - {device?.limit}
        </h4>
        <UpdateDeviceBtn id={id} />
      </div>

      <div className="m-10 flex justify-center">
        {device && (
          <TemperatureChart id={id} limit={device?.limit} className="w-full" />
        )}
      </div>
    </div>
  );
};

export default DevicePage;
