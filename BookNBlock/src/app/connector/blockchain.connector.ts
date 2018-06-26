import { Injectable } from '@angular/core';
import { Offer } from '../data/offer';
import { OpenDoorMessage } from '../data/OpenDoorMessage';
import { User } from '../data/user';

/**
 * Interface for all blockchains.
 */
@Injectable()
export abstract class BlockchainConnector {

  /**
   * Configure the connector to work with the user.
   * @param user The user.
   */
  abstract init(user: User): BlockchainConnector;

  /**
   * Get one offer from the blockchain.
   * @param id Offer Id.
   */
  abstract getOffer(id: number): Promise<Offer>;

  /**
   * Get all free offers from the blockchain.
   */
  abstract getAllOffers(from: Date, to: Date): Promise<Offer[]>;

  /**
   * Search for offers meeting one criterion.
   * @param criterion The search criterion.
   */
  abstract searchOffer(criterion: any): Promise<Offer[]>;

  /**
   * Add a new offer to the blockchain.
   * @param offer The new offer.
   */
  abstract insertOffer(offer: Offer): Promise<number>;

  /**
   * Rent an offer.
   * @param offerId: Id of the offer.
   * @param checkIn: Optional. Date of check-in, if none is provided the offer fromDate is used.
   * @param checkOut: Optional. Date of check-out, if none is provided the offer toDate is used.
   */
  abstract rentOffer(offerId: number, checkIn?: Date, checkOut?: Date): Promise<void>;

  /**
   * Send a message to open a door.
   * @param message The message.
   */
  abstract sendMessage(message: OpenDoorMessage): Promise<void>;

  /**
   * Check if a user as an account for the blockchain.
   * @param user The user information.
   */
  abstract authenticateUser(user: any): Promise<boolean>;

}
