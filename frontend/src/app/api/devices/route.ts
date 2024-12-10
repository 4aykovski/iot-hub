import { NextResponse } from "next/server";
import { getDevices } from "../../../../server/api";

export async function GET(req: NextResponse) {
  try {
    const res = await getDevices();

    if (!res) {
      return NextResponse.error();
    }

    const devices: Device[] = res.devices.map((device: any) => {
      return {
        ID: device.id,
        name: device.name,
        type: device.type,
        limit: device.limit,
      };
    });

    return NextResponse.json(devices);
  } catch (error: any) {
    console.error(error);
    return NextResponse.error();
  }
}
