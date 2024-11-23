import DeviceList from "@/components/DeviceList";

export default function Home() {
  const devices = [{ name: "Device 1", data: "Data 1", ID: "1" }];

  return (
    <div>
      <main>
        <DeviceList devices={devices} type="card" />
      </main>
    </div>
  );
}
