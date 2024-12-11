export const getDevices = async () => {
  const response = await fetch("http://localhost:8080/api/v1/devices");

  if (!response.ok) {
    throw new Error(response.statusText);
  }

  return response.json();
};

export const getDevice = async (id: string) => {
  const response = await fetch(`http://localhost:8080/api/v1/devices/${id}`);

  if (!response.ok) {
    throw new Error(response.statusText);
  }

  return response.json();
};

export const updateDevice = async (id: string, data: any) => {
  console.log(data);
  const response = await fetch(`http://localhost:8080/api/v1/devices/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });

  if (!response.ok) {
    throw new Error(response.statusText);
  }

  return response.json();
};

export const getData = async (id: string, interval: string) => {
  const response = await fetch(
    `http://localhost:8080/api/v1/devices/${id}/data?interval=${interval}`,
  );

  console.log(response);

  if (!response.ok) {
    throw new Error(response.statusText);
  }

  return response.json();
};
