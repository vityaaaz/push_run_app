import { ApiProperty } from '@nestjs/swagger';

export class ChallengeParticipant {
  @ApiProperty()
  userId: string;

  @ApiProperty()
  progress: number;

  @ApiProperty()
  isCompleted: boolean;

  @ApiProperty()
  completedAt?: Date;

  @ApiProperty()
  joinedAt: Date;
}

export class Challenge {
  @ApiProperty()
  id: string;

  @ApiProperty()
  name: string;

  @ApiProperty()
  description: string;

  @ApiProperty()
  type: string;

  @ApiProperty()
  target: number;

  @ApiProperty()
  startDate: Date;

  @ApiProperty()
  endDate: Date;

  @ApiProperty()
  createdBy: string;

  @ApiProperty({ type: [ChallengeParticipant] })
  participants: ChallengeParticipant[];

  @ApiProperty()
  createdAt: Date;

  @ApiProperty()
  updatedAt: Date;
}

export class CreateChallengeRequest {
  @ApiProperty()
  name: string;

  @ApiProperty()
  description: string;

  @ApiProperty()
  type: string;

  @ApiProperty()
  target: number;

  @ApiProperty()
  startDate: Date;

  @ApiProperty()
  endDate: Date;
}

export class JoinChallengeRequest {
  @ApiProperty()
  challengeId: string;
}

export class ChallengeProgress {
  @ApiProperty()
  current: number;

  @ApiProperty()
  target: number;

  @ApiProperty()
  percentage: number;

  @ApiProperty()
  daysLeft: number;
}

export class ChallengeWithProgress extends Challenge {
  @ApiProperty({ type: ChallengeProgress })
  progress: ChallengeProgress;

  @ApiProperty()
  isCompleted: boolean;

  @ApiProperty({ required: false })
  completedAt?: Date;
} 