import { NextResponse } from "next/server";
import { getDevice, updateDevice } from "../../../../../server/api";

export async function GET(
  req: Request,
  { params }: { params: { id: string } },
) {
  try {
    const p = await params;
    const id = p.id;
    if (!id) {
      return NextResponse.error();
    }

    const res = await getDevice(id);

    const device: Device = {
      ID: res.device.id,
      name: res.device.name,
      type: res.device.type,
      limit: res.device.limit,
    };

    return NextResponse.json(device);
  } catch (error: any) {
    console.error("123");
    console.error(error);
    return NextResponse.error();
  }
}

export async function PUT(
  req: Request,
  { params }: { params: { id: string } },
) {
  try {
    const p = await params;
    const id = p.id;
    if (!id) {
      return NextResponse.error();
    }

    const body = await req.json();
    const res = await updateDevice(id, body);

    return NextResponse.json(res);
  } catch (error: any) {
    console.log("123");
    console.error(error);
    return NextResponse.error();
  }
}

export async function POST(
  req: Request,
  { params }: { params: { id: string } },
) {
  try {
    const p = await params;
    const id = p.id;
    if (!id) {
      return NextResponse.error();
    }

    const body = await req.json();
    const res = await updateDevice(id, body);

    return NextResponse.json(res);
  } catch (error: any) {
    console.error(error);
    return NextResponse.error();
  }
}
