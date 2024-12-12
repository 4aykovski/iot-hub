import React from "react";

interface DataListProps {
  id: string;
}

const DataList: React.FC<DataListProps> = ({ id }) => {
  const [chartData, setChartData] = React.useState<Data[]>([]);

  React.useEffect(() => {
    const fetchData = async () => {
      const response = await fetch(`/api/devices/${id}/data`);
      const data = await response.json();
      setChartData(data);
    };
    fetchData();
    const intervalId = setInterval(fetchData, 5000);

    return () => clearInterval(intervalId);
  }, [id]);

  return (
    <div>
      <table className="table-auto">
        <thead>
          <tr>
            <th className="px-4 py-2">Значение</th>
            <th className="px-4 py-2">Время</th>
          </tr>
        </thead>
        <tbody>
          {chartData.map((data) => (
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
