import React from "react";

interface DataListProps {
  chartData: Data[];
  className?: string;
}

const DataList: React.FC<DataListProps> = ({ chartData, className }) => {
  const reversedChartData = [...chartData].reverse();
  return (
    <div className={className}>
      <table className="table-auto">
        <thead>
          <tr>
            <th className="px-4 py-2">Значение</th>
            <th className="px-4 py-2">Время</th>
          </tr>
        </thead>
        <tbody>
          {reversedChartData.map((data) => (
            <tr key={data.ID}>
              <td className="border px-4 py-2">{data.value}</td>
              <td className="border px-4 py-2">{data.timestamp.toString()}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default DataList;
