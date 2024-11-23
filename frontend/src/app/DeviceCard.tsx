import TemperatuteChart from "@/components/TemperatureChart";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import Link from "next/link";

interface DeviceCardProps {
  device: Device;
}

const DeviceCard: React.FC<DeviceCardProps> = ({ device }) => {
  return (
    <Card className="bg-primary-foreground w-[400px] h-auto  mb-10">
      <CardHeader>
        <CardTitle>
          <Link href={`/devices/${device.ID}`} legacyBehavior passHref>
            {device.name}
          </Link>
        </CardTitle>
      </CardHeader>
      <CardContent>
        <TemperatuteChart id={device.ID} />
      </CardContent>
    </Card>
  );
};

export default DeviceCard;
