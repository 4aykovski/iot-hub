"use client";

import { Area, AreaChart, CartesianGrid, LabelList, XAxis } from "recharts";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { cn } from "@/lib/utils";
import React from "react";
import IntervalSelector from "./IntervalSelector";
import { toast } from "sonner";

const chartConfig = {
  temperature: {
    label: "Temperature",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig;

interface TemperatureChartProps {
  id: string;
  limit: number;
  className?: string;
}

const TemperatuteChart: React.FC<TemperatureChartProps> = ({
  id,
  limit,
  className,
}) => {
  const [lastNotified, setLastNotified] = React.useState<number>(0);
  const [chartData, setChartData] = React.useState<Data[]>([]);
  const [timeInterval, setTimeInterval] = React.useState("60");

  function handleIntervalSelect(interval) {
    setTimeInterval(interval);
  }

  React.useEffect(() => {
    const fetchData = async () => {
      const response = await fetch(
        `/api/devices/${id}/data?interval=${timeInterval}`,
      );
      const data = await response.json();
      setChartData(data);
    };
    fetchData();
    const intervalId = setInterval(fetchData, 5000);

    return () => clearInterval(intervalId);
  }, [id, timeInterval]);

  React.useEffect(() => {
    if (limit === -1) {
      return;
    }

    for (let i = 0; i < chartData.length; i++) {
      if (chartData[i].value > limit && lastNotified < chartData[i].ID) {
        toast(`Лимит превышен у устройства ${id}`, {
          description: `${chartData[i].value} > ${limit}, ${chartData[i].timestamp}`,
          action: {
            label: "Undo",
            onClick: () => console.log("Undo"),
          },
        });

        setLastNotified(chartData[i].ID);
      }
    }
  }, [chartData, limit, id]);

  return (
    <Card className={cn("max-w-[750px] min-w-[160px]", className)}>
      <CardHeader>
        <CardTitle>
          <IntervalSelector onChange={handleIntervalSelect} />
        </CardTitle>
        <CardDescription></CardDescription>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig} className="max-h-[400px]">
          <AreaChart
            accessibilityLayer
            data={chartData}
            margin={{
              left: 12,
              right: 12,
            }}
          >
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="timestamp"
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              tickFormatter={(value) => value.slice(0, 8)}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dot" />}
            />
            <Area
              dataKey="value"
              type="natural"
              fill="var(--color-temperature)"
              fillOpacity={0.4}
              dot={{
                fill: "var(--color-temperature)",
              }}
              activeDot={{
                r: 6,
              }}
              stroke="var(--color-temperature)"
              stackId="a"
            >
              <LabelList
                position="top"
                offset={12}
                className="fill-foreground"
                fontSize={6}
              />
            </Area>
          </AreaChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
};

export default TemperatuteChart;
