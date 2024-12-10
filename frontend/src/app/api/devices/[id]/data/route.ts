import { NextRequest, NextResponse } from "next/server";
import { getData } from "../../../../../../server/api";

export async function GET(
  req: NextRequest,
  { params }: { params: { id: string } },
) {
  try {
    const p = await params;
    const id = p.id;
    if (!id) {
      return NextResponse.error();
    }

    const searchParams = req.nextUrl.searchParams;
    const interval = searchParams.get("interval") || "30";

    const res = await getData(id, interval);

    if (!res) {
      return NextResponse.error();
    }
    console.log(res);

    const data: Data[] = res.data.map((data: any) => {
      return {
        ID: data.ID,
        value: data.Value,
        timestamp: new Date(data.Timestamp).toTimeString(),
      };
    });

    console.log(data);

    return NextResponse.json(data);
  } catch (error: any) {
    console.error(error);
    return NextResponse.error();
  }
}
