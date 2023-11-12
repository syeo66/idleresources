import { z } from "zod"

export const resourceIdSchema = z.enum(["water"])
export type ResourceId = z.infer<typeof resourceIdSchema>

export const resourceSchema = z.object({
  id: resourceIdSchema, 
  name: z.string(),
  amount: z.number(),
  delta: z.number(),
  total: z.number(),
  is_automated: z.boolean(),
})
export type Resource = z.infer<typeof resourceSchema>

export const gameStateSchema = z.object({
  resources: z.array(resourceSchema)
})
export type GameState = z.infer<typeof gameStateSchema>
