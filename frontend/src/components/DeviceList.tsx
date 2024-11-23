import DeviceCard from "@/app/DeviceCard";
import DeviceRow from "@/app/devices/DeviceRow";

interface DeviceListProps {
  devices: Device[];
  type: string;
}

const DeviceList: React.FC<DeviceListProps> = ({ devices, type }) => {
  return type === "card" ? (
    <div className="flex justify-around flex-wrap">
      {devices.map((device) => (
        <DeviceCard key={device.ID} device={device} />
      ))}
    </div>
  ) : (
    <div>
      <table>
        <thead>
          <tr>
            <th className="w-[100px]">ID</th>
            <th className="w-[100px]">Name</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          {devices.map((device) => (
            <DeviceRow key={device.ID} device={device} />
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default DeviceList;
