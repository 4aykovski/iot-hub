import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

interface IntervalSelectorProps {
  placeholder?: string;
  onChange?: (value: string) => void;
}

const intervals = [
  { id: 1, name: "30s", interval: "30" },
  { id: 2, name: "60s", interval: "60" },
  { id: 3, name: "5m", interval: "300" },
  { id: 4, name: "10m", interval: "600" },
];

const IntervalSelector: React.FC<IntervalSelectorProps> = ({ onChange }) => {
  return (
    <Select onValueChange={onChange}>
      <SelectTrigger className="w-[260px]">
        <SelectValue placeholder="Выберите временной промежуток" />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          <SelectLabel></SelectLabel>
          {intervals?.map((item) => (
            <SelectItem key={item.id} value={item.interval}>
              {item.name}
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  );
};

export default IntervalSelector;
