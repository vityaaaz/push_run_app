import { ApiProperty } from '@nestjs/swagger';

export class Coordinate {
  @ApiProperty()
  latitude: number;

  @ApiProperty()
  longitude: number;

  @ApiProperty()
  altitude: number;

  @ApiProperty()
  timestamp: Date;
}

export class HeartRate {
  @ApiProperty()
  value: number;

  @ApiProperty()
  timestamp: Date;
}

export class Weather {
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
}

export class Comment {
  @ApiProperty()
  id: string;

  @ApiProperty()
  userId: string;

  @ApiProperty()
  text: string;

  @ApiProperty()
  createdAt: Date;
}

export class Activity {
  @ApiProperty()
  id: string;

  @ApiProperty()
  userId: string;

  @ApiProperty()
  type: string;

  @ApiProperty()
  startTime: Date;

  @ApiProperty()
  endTime: Date;

  @ApiProperty()
  distance: number;

  @ApiProperty()
  duration: number;

  @ApiProperty()
  pace: number;

  @ApiProperty()
  calories: number;

  @ApiProperty()
  elevationGain: number;

  @ApiProperty({ type: [Coordinate] })
  coordinates: Coordinate[];

  @ApiProperty({ type: [HeartRate], required: false })
  heartRate?: HeartRate[];

  @ApiProperty({ type: Weather, required: false })
  weather?: Weather;

  @ApiProperty({ required: false })
  notes?: string;

  @ApiProperty()
  isPublic: boolean;

  @ApiProperty({ type: [String] })
  likes: string[];

  @ApiProperty({ type: [Comment] })
  comments: Comment[];

  @ApiProperty()
  createdAt: Date;

  @ApiProperty()
  updatedAt: Date;
}

export class CreateActivityRequest {
  @ApiProperty()
  type: string;

  @ApiProperty()
  startTime: Date;

  @ApiProperty()
  endTime: Date;

  @ApiProperty()
  distance: number;

  @ApiProperty({ type: [Coordinate] })
  coordinates: Coordinate[];

  @ApiProperty({ required: false })
  calories?: number;

  @ApiProperty({ required: false })
  elevationGain?: number;

  @ApiProperty({ type: [HeartRate], required: false })
  heartRate?: HeartRate[];

  @ApiProperty({ type: Weather, required: false })
  weather?: Weather;

  @ApiProperty({ required: false })
  notes?: string;

  @ApiProperty()
  isPublic: boolean;
}

export class ActivityStats {
  @ApiProperty()
  totalDistance: number;

  @ApiProperty()
  totalTime: number;

  @ApiProperty()
  totalActivities: number;

  @ApiProperty()
  averagePace: number;

  @ApiProperty()
  bestPace: number;

  @ApiProperty()
  totalElevationGain: number;

  @ApiProperty()
  totalCalories: number;
} 