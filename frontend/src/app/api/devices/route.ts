import { NextResponse } from "next/server";
import { getDevices } from "../../../../server/api";

export async function GET(req: NextResponse) {
  try {
    const res = await getDevices();

    return NextResponse.json(res);
  } catch (error: any) {
    console.error(error);
    return NextResponse.error();
  }
}
