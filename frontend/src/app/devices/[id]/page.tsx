"use client";

import TemperatureChart from "@/components/TemperatureChart";
import { useParams } from "next/navigation";
import React from "react";
import UpdateDeviceBtn from "./UpdateDeviceBtn";
import DataList from "./DataList";
import { Button } from "@/components/ui/button";

const DevicePage: React.FC = () => {
  const [device, setDevice] = React.useState<Device | null>(null);
  const [deviceData, setDeviceData] = React.useState<Data[]>([]);
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

  function handleBtnClick() {
    const fetchDeviceData = async () => {
      const response = await fetch(`/api/devices/${id}/data?interval=-1`);
      const data = await response.json();
      console.log(data);
      setDeviceData(data);
    };
    fetchDeviceData();
  }

  return (
    <div className="mb-10 flex flex-col justify-center">
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

      <div className="flex justify-center items-center flex-col">
        <div className="overflow-hidden w-[500px] h-[500px]">
          <Button variant="outline" onClick={handleBtnClick}>
            Показать данные
          </Button>
          <DataList
            className="overflow-y-scroll h-[450px]"
            chartData={deviceData}
          />
        </div>
      </div>
    </div>
  );
};

export default DevicePage;
