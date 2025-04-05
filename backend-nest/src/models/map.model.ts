import { ApiProperty } from '@nestjs/swagger';

export class RoutePoint {
  @ApiProperty()
  latitude: number;

  @ApiProperty()
  longitude: number;

  @ApiProperty()
  altitude: number;
}

export class Route {
  @ApiProperty()
  id: string;

  @ApiProperty()
  name: string;

  @ApiProperty()
  description: string;

  @ApiProperty({ type: [RoutePoint] })
  points: RoutePoint[];

  @ApiProperty()
  distance: number;

  @ApiProperty()
  elevationGain: number;

  @ApiProperty()
  createdBy: string;

  @ApiProperty()
  isPublic: boolean;

  @ApiProperty()
  createdAt: Date;

  @ApiProperty()
  updatedAt: Date;
}

export class CreateRouteRequest {
  @ApiProperty()
  name: string;

  @ApiProperty()
  description: string;

  @ApiProperty({ type: [RoutePoint] })
  points: RoutePoint[];

  @ApiProperty()
  isPublic: boolean;
}

export class WeatherInfo {
  @ApiProperty()
  temperature: number;

  @ApiProperty()
  feelsLike: number;

  @ApiProperty()
  condition: string;

  @ApiProperty()
  windSpeed: number;

  @ApiProperty()
  humidity: number;

  @ApiProperty()
  timestamp: Date;
}

export class MapBounds {
  @ApiProperty()
  north: number;

  @ApiProperty()
  south: number;

  @ApiProperty()
  east: number;

  @ApiProperty()
  west: number;
} 