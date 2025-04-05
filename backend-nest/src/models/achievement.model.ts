import { ApiProperty } from '@nestjs/swagger';

export class Achievement {
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
  icon: string;

  @ApiProperty()
  createdAt: Date;
}

export class UserAchievement {
  @ApiProperty()
  id: string;

  @ApiProperty()
  userId: string;

  @ApiProperty()
  achievementId: string;

  @ApiProperty()
  progress: number;

  @ApiProperty()
  isCompleted: boolean;

  @ApiProperty()
  completedAt?: Date;

  @ApiProperty()
  createdAt: Date;

  @ApiProperty()
  updatedAt: Date;
}

export class AchievementProgress {
  @ApiProperty()
  current: number;

  @ApiProperty()
  target: number;

  @ApiProperty()
  percentage: number;
}

export class AchievementWithProgress extends Achievement {
  @ApiProperty({ type: AchievementProgress })
  progress: AchievementProgress;

  @ApiProperty()
  isCompleted: boolean;

  @ApiProperty({ required: false })
  completedAt?: Date;
} 