"use client";

import { Area, AreaChart, CartesianGrid, XAxis } from "recharts";
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

const chartData = [
  { time: "01:00", temperature: "10" },
  { time: "02:00", temperature: "5" },
  { time: "03:00", temperature: "6" },
  { time: "04:00", temperature: "9" },
  { time: "05:00", temperature: "11" },
];

const chartConfig = {
  temperature: {
    label: "Temperature",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig;

interface TemperatureChartProps {
  id: string;
  className?: string;
}

const TemperatuteChart: React.FC<TemperatureChartProps> = ({
  id,
  className,
}) => {
  return (
    <Card className={cn("max-w-[750px] min-w-[160px]", className)}>
      <CardHeader>
        <CardTitle>Температура</CardTitle>
        <CardDescription>Температура за последний час</CardDescription>
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
              dataKey="time"
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              tickFormatter={(value) => value.slice(1, 5)}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dot" />}
            />
            <Area
              dataKey="temperature"
              type="natural"
              fill="var(--color-temperature)"
              fillOpacity={0.4}
              stroke="var(--color-temperature)"
              stackId="a"
            />
          </AreaChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
};

export default TemperatuteChart;
