import TemperatuteChart from "@/components/TemperatureChart";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import Link from "next/link";

interface DeviceCardProps {
  device: Device;
}

const DeviceCard: React.FC<DeviceCardProps> = ({ device }) => {
  console.log(device);
  return (
    <Card className="bg-primary-foreground w-[400px] h-auto  mb-10">
      <CardHeader className="flex justify-between flex-row">
        <CardTitle>
          <Link href={`/devices/${device.ID}`} legacyBehavior passHref>
            {device.name}
          </Link>
        </CardTitle>
        <div>ID: {device.ID}</div>
      </CardHeader>
      <CardContent>
        <TemperatuteChart id={device.ID} limit={device.limit} />
      </CardContent>
    </Card>
  );
};

export default DeviceCard;
