import { NextResponse } from "next/server";
import { getData } from "../../../../../../server/api";

export async function GET(req: NextResponse) {
  try {
    const { searchParams } = new URL(req.url);
    const id = searchParams.get("id");
    if (!id) {
      return NextResponse.error();
    }

    const res = await getData(id);

    return NextResponse.json(res);
  } catch (error: any) {
    console.error(error);
    return NextResponse.error();
  }
}
