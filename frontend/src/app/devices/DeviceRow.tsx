import { Button } from "@/components/ui/button";
import Link from "next/link";

interface DeviceRowProps {
  device: Device;
}

const DeviceRow: React.FC<DeviceRowProps> = ({ device }) => {
  return (
    <tr>
      <td className="text-center">{device.ID}</td>
      <td className="text-center">{device.name}</td>
      <td>
        <Button variant="link">
          <Link href={`/devices/${device.ID}`}>Details</Link>
        </Button>
      </td>
    </tr>
  );
};

export default DeviceRow;
