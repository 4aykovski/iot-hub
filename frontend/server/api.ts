export const getDevices = async () => {
  const response = await fetch("/api/v1/devices");

  if (!response.ok) {
    throw new Error(response.statusText);
  }

  return response.json();
};

export const getDevice = async (id: string) => {
  const response = await fetch(`/api/v1/devices/${id}`);

  if (!response.ok) {
    throw new Error(response.statusText);
  }

  return response.json();
};

export const updateDevice = async (id: string, data: any) => {
  const response = await fetch(`/api/v1/devices/${id}`, {
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

export const getData = async (id: string) => {
  const response = await fetch(`/api/v1/devices/${id}/data`);

  if (!response.ok) {
    throw new Error(response.statusText);
  }

  return response.json();
};
