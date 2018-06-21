export class Offer {
  id: number;
  doorId: string;
  prize: number;
  fromDate: Date;
  toDate: Date;
  address: string;
  nameLandlord: string;
  description: string;
  walletId: string;
  title: string;
  image?: any;

  constructor(id: number,
    doorId: number,
    isBooked: boolean,
    prize: number,
    fromDate: Date,
    toDate: Date,
    address: string,
    name: string,
    description: string,
    walletId: number,
    title: string) {
    this.id = id;
    this.doorId = doorId;
    this.isBooked = isBooked;
    this.prize = prize;
    this.fromDate = fromDate;
    this.toDate = toDate;
    this.address = address;
    this.name = name;
    this.description = description;
    this.walletId = walletId;
    this.title = title;
    this.image = undefined;
  }
}
