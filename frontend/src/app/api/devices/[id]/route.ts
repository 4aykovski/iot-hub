import { NextResponse } from "next/server";
import { getDevice, updateDevice } from "../../../../../server/api";

export async function GET(req: Request) {
  try {
    const { searchParams } = new URL(req.url);
    const id = searchParams.get("id");
    if (!id) {
      return NextResponse.error();
    }

    const res = await getDevice(id);

    return NextResponse.json(res);
  } catch (error: any) {
    console.error(error);
    return NextResponse.error();
  }
}

export async function PUT(req: Request) {
  try {
    const { searchParams } = new URL(req.url);
    const id = searchParams.get("id");
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
