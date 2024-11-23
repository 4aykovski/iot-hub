import TemperatureChart from "@/components/TemperatureChart";

interface DevicePageProps {
  params: Promise<{ id: string }>;
}

const DevicePage: React.FC<DevicePageProps> = async ({ params }) => {
  const id = (await params).id;

  return (
    <div className="mb-10 flex flex-col">
      <h4 className="mx-14 scroll-m-20 text-xl font-semibold tracking-tight">
        Device {id} - name
      </h4>

      <div className="m-10 flex justify-center">
        <TemperatureChart id={id} className="w-full" />
      </div>
    </div>
  );
};

export default DevicePage;
